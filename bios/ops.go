package bios

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
	ResetTestnetOptions() // TODO: implement the DISABLING of all testnet options when `mainnet` is voted in the `discovery`.
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

func (op *OpSetCode) ResetTestnetOptions() { return }
func (op *OpSetCode) Actions(b *BIOS) ([]*eos.Action, error) {
	wasmFileRef, err := b.GetContentsCacheRef(fmt.Sprintf("%s.wasm", op.ContractNameRef))
	if err != nil {
		return nil, err
	}
	abiFileRef, err := b.GetContentsCacheRef(fmt.Sprintf("%s.abi", op.ContractNameRef))
	if err != nil {
		return nil, err
	}

	setCode, err := system.NewSetCodeTx(
		op.Account,
		b.Network.FileNameFromCache(wasmFileRef),
		b.Network.FileNameFromCache(abiFileRef),
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

func (op *OpNewAccount) ResetTestnetOptions() { return }
func (op *OpNewAccount) Actions(b *BIOS) (out []*eos.Action, err error) {
	pubKey := *b.EphemeralPublicKey
	if op.Pubkey != "ephemeral" {
		pubKey, err = ecc.NewPublicKey(op.Pubkey)
		if err != nil {
			return nil, fmt.Errorf("reading pubkey: %s", err)
		}
	}

	return append(out, system.NewNewAccount(op.Creator, op.NewAccount, pubKey)), nil
}

//

type OpSetPriv struct {
	Account eos.AccountName
	IsPriv  bool `json:"is_priv"` // unused
}

func (op *OpSetPriv) ResetTestnetOptions() { return }
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

func (op *OpCreateToken) ResetTestnetOptions() {}
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

func (op *OpIssueToken) ResetTestnetOptions() {}
func (op *OpIssueToken) Actions(b *BIOS) (out []*eos.Action, err error) {
	act := token.NewIssue(op.Account, op.Amount, op.Memo)
	return append(out, act), nil
}

//

type OpCreateProducers struct {
	// TestnetEnrichProducers will provide each producer account with some EOS, only on testnets.
	TestnetEnrichProducers bool `json:"TESTNET_ENRICH_PRODUCERS"`
}

func (op *OpCreateProducers) ResetTestnetOptions() {
	op.TestnetEnrichProducers = false
}

func (op *OpCreateProducers) Actions(b *BIOS) (out []*eos.Action, err error) {
	for _, prod := range b.ShuffledProducers {
		newAccount := system.NewNewAccount(AN("eosio"), AN(prod.AccountName()), nil)
		newAccount.Data = eos.NewActionData(system.NewAccount{
			Creator:  AN("eosio"),
			Name:     AN(prod.AccountName()),
			Owner:    prod.Discovery.TargetInitialAuthority.Owner,
			Active:   prod.Discovery.TargetInitialAuthority.Active,
			Recovery: prod.Discovery.TargetInitialAuthority.Recovery,
		})

		// mama, _ := json.MarshalIndent(newAccount.Data, "", "  ")
		// fmt.Println("Some JSON", string(mama))

		fmt.Printf("- Creating new account %q\n", prod.AccountName())
		out = append(out, newAccount)

		if op.TestnetEnrichProducers {
			fmt.Printf("  DEBUG: Enriching producer %q\n", prod.AccountName())
			out = append(out, token.NewTransfer(AN("eosio"), AN(prod.AccountName()), eos.NewEOSAsset(1000000000), "Hey, make good use of it!"))
		}
	}
	return
}

//

type OpInjectSnapshot struct {
	TestnetTruncateSnapshot int `json:"TESTNET_TRUNCATE_SNAPSHOT"`
}

func (op *OpInjectSnapshot) ResetTestnetOptions() {
	op.TestnetTruncateSnapshot = 0
}

func (op *OpInjectSnapshot) Actions(b *BIOS) (out []*eos.Action, err error) {
	snapshotFile, err := b.GetContentsCacheRef("snapshot.csv")
	if err != nil {
		return nil, err
	}

	rawSnapshot, err := b.Network.ReadFromCache(snapshotFile)
	if err != nil {
		return nil, fmt.Errorf("reading snapshot file: %s", err)
	}

	snapshotData, err := NewSnapshot(rawSnapshot)
	if err != nil {
		return nil, fmt.Errorf("loading snapshot csv: %s", err)
	}

	if len(snapshotData) == 0 {
		return nil, fmt.Errorf("snapshot is empty or not loaded")
	}

	for idx, hodler := range snapshotData {
		flipped := flipEndianness(uint64(idx + 1))
		destAccount := AN("genesis." + eos.NameToString(flipped))
		fmt.Println("Transfer", hodler, destAccount)

		out = append(out, system.NewNewAccount(AN("eosio"), destAccount, hodler.EOSPublicKey))

		memo := "Welcome " + hodler.EthereumAddress[len(hodler.EthereumAddress)-6:]

		out = append(out, token.NewTransfer(AN("eosio"), destAccount, hodler.Balance, memo))

		if trunc := op.TestnetTruncateSnapshot; trunc != 0 {
			if idx == trunc {
				fmt.Printf("- DEBUG: truncated snapshot at %d rows\n", trunc)
				break
			}
		}

		// TODO: stake 50% bandwidth, 50% cpu for all new accounts
		// out = append(out, system.Stake(AN("eosio"), destAccount, 999, 888, ""))
	}

	return
}

//

type OpSetProds struct{}

func (op *OpSetProds) ResetTestnetOptions() {}
func (op *OpSetProds) Actions(b *BIOS) (out []*eos.Action, err error) {
	prodkeys := []system.ProducerKey{system.ProducerKey{
		ProducerName:    AN("eosio"),
		BlockSigningKey: b.EphemeralPrivateKey.PublicKey(),
	}}
	// prodkeys := []system.ProducerKey{}
	for _, prod := range b.ShuffledProducers {
		prodkeys = append(prodkeys, system.ProducerKey{AN(prod.AccountName()), prod.Discovery.TargetAppointedBlockProducerSigningKey})
		if len(prodkeys) >= 21 {
			break
		}
	}
	out = append(out, system.NewSetProds(0, prodkeys))

	return
}

//

type OpDestroyAccounts struct {
	Accounts            []eos.AccountName
	TestnetKeepAccounts bool `json:"TESTNET_KEEP_ACCOUNTS"`
}

func (op *OpDestroyAccounts) ResetTestnetOptions() {
	op.TestnetKeepAccounts = false
}

func (op *OpDestroyAccounts) Actions(b *BIOS) (out []*eos.Action, err error) {
	if op.TestnetKeepAccounts {
		fmt.Println("DEBUG: Keeping system account around, for testing purposes.")
		return
	}

	for _, acct := range op.Accounts {
		out = append(out,
			system.NewUpdateAuth(acct, PN("active"), PN("owner"), eos.Authority{
				Threshold: 1,
				Keys: []eos.KeyWeight{
					{
						PublicKey: ecc.PublicKey(make([]byte, 34, 34)),
						Weight:    1,
					},
				},
			}, PN("active")),
			system.NewUpdateAuth(acct, PN("owner"), PN(""), eos.Authority{
				Threshold: 1,
				Keys: []eos.KeyWeight{
					{
						PublicKey: ecc.PublicKey(make([]byte, 34, 34)),
						Weight:    1,
					},
				},
			}, PN("owner")),
			// TODO: add recovery here ??
		)

		// unregister the producer at the same time ?
	}
	return
}
