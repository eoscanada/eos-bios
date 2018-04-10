package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
	"github.com/eoscanada/eos-go/system"
	"github.com/eoscanada/eos-go/token"
)

type OperationType struct {
	Op    string
	Label string
	Data  Operation
}

func (o *OperationType) UnmarshalJSON(data []byte) error {
	opData := struct {
		Op    string
		Label string
		Data  json.RawMessage
	}{}
	if err := json.Unmarshal(data, &opData); err != nil {
		return err
	}

	opType, found := operationsRegistry[opData.Op]
	if !found {
		return fmt.Errorf("operation type %q invalid, use one of: %q", opData.Op, operationsRegistry)
	}

	objType := reflect.TypeOf(opType).Elem()
	obj := reflect.New(objType).Interface()

	if len(opData.Data) != 0 {
		err := json.Unmarshal(opData.Data, &obj)
		if err != nil {
			return fmt.Errorf("operation type %q invalid, error decoding: %s", opData.Op, err)
		}
	} //  else {
	// 	_ = json.Unmarshal([]byte("{}"), &obj)
	// }

	opIface, ok := obj.(Operation)
	if !ok {
		return fmt.Errorf("operation type %q isn't an op", opData.Op)
	}

	*o = OperationType{
		Op:    opData.Op,
		Label: opData.Label,
		Data:  opIface,
	}

	return nil
}

type Operation interface {
	Actions(b *BIOS) ([]*eos.Action, error)
}

var operationsRegistry = map[string]Operation{
	"system.setcode":            &OpSetCode{},
	"system.newaccount":         &OpNewAccount{},
	"system.setpriv":            &OpSetPriv{},
	"token.create":              &OpCreateToken{},
	"token.issue":               &OpIssueToken{},
	"producers.create_accounts": &OpCreateProducers{},
	"system.setprods":           &OpSetProds{},
	"snapshot.inject":           &OpInjectSnapshot{},
	"system.destroy_accounts":   &OpDestroyAccounts{},
}

//

type OpSetCode struct {
	Account         eos.AccountName
	ContractNameRef string `json:"contract_name_ref"`
}

func (op OpSetCode) Actions(b *BIOS) ([]*eos.Action, error) {
	setCode, err := system.NewSetCodeTx(
		op.Account,
		b.Config.Contracts[op.ContractNameRef].CodePath,
		b.Config.Contracts[op.ContractNameRef].ABIPath,
	)
	if err != nil {
		return nil, fmt.Errorf("NewSetCodeTx %s: %s", op.ContractNameRef, err)
	}

	return setCode.Actions, nil
}

//

type OpNewAccount struct {
	Creator    eos.AccountName
	NewAccount eos.AccountName `json:"new_account"`
	Pubkey     string
}

func (op OpNewAccount) Actions(b *BIOS) (out []*eos.Action, err error) {
	pubkey := b.EphemeralPrivateKey.PublicKey()
	if op.Pubkey != "ephemeral" {
		pubkey, err = ecc.NewPublicKey(op.Pubkey)
		if err != nil {
			return nil, fmt.Errorf("reading pubkey: %s", err)
		}
	}

	return append(out, system.NewNewAccount(op.Creator, op.NewAccount, pubkey)), nil
}

//

type OpSetPriv struct {
	Account eos.AccountName
	IsPriv  bool `json:"is_priv"` // unused
}

func (op *OpSetPriv) Actions(b *BIOS) (out []*eos.Action, err error) {
	return append(out, system.NewSetPriv(op.Account)), nil
}

//

type OpCreateToken struct {
	Account      eos.AccountName
	Amount       eos.Asset
	CanWhitelist bool `json:"can_whitelist"`
	CanFreeze    bool `json:"can_freeze"`
	CanRecall    bool `json:"can_recall"`
}

func (op *OpCreateToken) Actions(b *BIOS) (out []*eos.Action, err error) {
	act := token.NewCreate(op.Account, op.Amount, op.CanFreeze, op.CanRecall, op.CanWhitelist)
	return append(out, act), nil
}

//

type OpIssueToken struct {
	Account eos.AccountName
	Amount  eos.Asset
	Memo    string
}

func (op *OpIssueToken) Actions(b *BIOS) (out []*eos.Action, err error) {
	act := token.NewIssue(op.Account, op.Amount, op.Memo)
	return append(out, act), nil
}

//

type OpCreateProducers struct{}

func (op *OpCreateProducers) Actions(b *BIOS) (out []*eos.Action, err error) {
	for _, prod := range b.ShuffledProducers {
		newAccount := system.NewNewAccount(AN("eosio"), prod.AccountName, nil)
		newAccount.Data = eos.NewActionData(system.NewAccount{
			Creator: AN("eosio"),
			Name:    prod.AccountName,
			Owner:   prod.Authority.Owner,
			Active:  prod.Authority.Active,
		})

		fmt.Printf("- Creating new account %q\n", prod.AccountName)
		out = append(out, newAccount)

		if b.Config.Debug.EnrichProducers {
			fmt.Printf("  DEBUG: Enriching producer %q\n", prod.AccountName)
			out = append(out, token.NewTransfer(AN("eosio"), prod.AccountName, eos.NewEOSAsset(1000000000), "Hey, make good use of it!"))
		}
	}
	return
}

//

type OpInjectSnapshot struct{}

func (op *OpInjectSnapshot) Actions(b *BIOS) (out []*eos.Action, err error) {
	for idx, hodler := range b.Snapshot {
		flipped := flipEndianness(uint64(idx + 1))
		destAccount := AN("genesis." + eos.NameToString(flipped))
		fmt.Println("Transfer", hodler, destAccount)

		out = append(out, system.NewNewAccount(AN("eosio"), destAccount, hodler.EOSPublicKey))

		memo := "Welcome " + hodler.EthereumAddress[len(hodler.EthereumAddress)-6:]

		out = append(out, token.NewTransfer(AN("eosio"), destAccount, hodler.Balance, memo))

		// 400 transfers per Tx before hitting some limits..
		// FIXME: this should be in the wrapping Tx mapper..
		if idx == 5000 {
			fmt.Println("- skipping remaining snapshot transfers")
			break
		}

		// TODO: stake 50% bandwidth, 50% cpu for all new accounts
		// b.API.SignPushActions(system.Stake(AN("eosio"), destAccount, 999, 888, ""))
	}

	return
}

//

type OpSetProds struct{}

func (op *OpSetProds) Actions(b *BIOS) (out []*eos.Action, err error) {
	// FIXME: shouldn't include eosio
	prodkeys := []system.ProducerKey{system.ProducerKey{
		ProducerName:    AN("eosio"),
		BlockSigningKey: b.EphemeralPrivateKey.PublicKey(),
	}}
	for idx, prod := range b.ShuffledProducers {
		if idx >= 20 { // FIXME: should be tweaked to make it a max of 21
			break
		}
		prodkeys = append(prodkeys, system.ProducerKey{prod.AccountName, prod.InitialBlockSigningPublicKey})
	}
	out = append(out, system.NewSetProds(0, prodkeys))

	return
}

//

type OpDestroyAccounts struct {
	Accounts []eos.AccountName
}

func (op *OpDestroyAccounts) Actions(b *BIOS) (out []*eos.Action, err error) {
	if b.Config.Debug.KeepSystemAccount {
		fmt.Println("DEBUG: Keeping system account around, for testing purposes.")
		return
	}

	for _, acct := range op.Accounts {
		out = append(out,
			system.NewUpdateAuth(acct, PN("active"), PN("owner"), eos.Authority{Threshold: 0}, PN("active")),
			system.NewUpdateAuth(acct, PN("owner"), PN(""), eos.Authority{Threshold: 0}, PN("owner")),
		)
	}
	return
}
