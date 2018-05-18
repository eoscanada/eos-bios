package bios

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
	"github.com/eoscanada/eos-go/system"
	"github.com/eoscanada/eos-go/token"
)

type Operation interface {
	Actions(b *BIOS) ([]*eos.Action, error)
	ResetTestnetOptions() // TODO: implement the DISABLING of all testnet options when `mainnet` is voted in the `discovery`.
}

var operationsRegistry = map[string]Operation{
	"system.setcode":            &OpSetCode{},
	"system.setram":             &OpSetRAM{},
	"system.newaccount":         &OpNewAccount{},
	"system.setpriv":            &OpSetPriv{},
	"token.create":              &OpCreateToken{},
	"token.issue":               &OpIssueToken{},
	"producers.create_accounts": &OpCreateProducers{},
	"producers.register":        &OpRegisterProducers{},
	"system.setprods":           &OpSetProds{},
	"snapshot.inject":           &OpInjectSnapshot{},
	"system.destroy_accounts":   &OpDestroyAccounts{},
}

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

type OpSetRAM struct {
	MaxRAMSize uint64 `json:"max_ram_size"`
}

func (op *OpSetRAM) ResetTestnetOptions() { return }
func (op *OpSetRAM) Actions(b *BIOS) (out []*eos.Action, err error) {
	return append(out, system.NewSetRAM(op.MaxRAMSize)), nil
}

//

type OpNewAccount struct {
	Creator    eos.AccountName
	NewAccount eos.AccountName `json:"new_account"`
	Pubkey     string
}

func (op *OpNewAccount) ResetTestnetOptions() { return }
func (op *OpNewAccount) Actions(b *BIOS) (out []*eos.Action, err error) {
	pubKey := b.EphemeralPublicKey
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
		prodName := prod.Discovery.TargetAccountName
		if prodName == AN("eosio") {
			prodName = prod.Discovery.SeedNetworkAccountName // only happens with --single
		}

		newAccount := system.NewNewAccount(AN("eosio"), prodName, ecc.PublicKey{}) // overridden just below
		newAccount.ActionData = eos.NewActionData(system.NewAccount{
			Creator: AN("eosio"),
			Name:    prodName,
			Owner:   prod.Discovery.TargetInitialAuthority.Owner,
			Active:  prod.Discovery.TargetInitialAuthority.Active,
		})
		buyRAMBytes := system.NewBuyRAMBytes(AN("eosio"), prodName, 8192) // 8kb gift !
		delegateBW := system.NewDelegateBW(AN("eosio"), prodName, eos.NewEOSAsset(10000), eos.NewEOSAsset(10000), true)

		// mama, _ := json.MarshalIndent(newAccount.Data, "", "  ")
		// fmt.Println("Some JSON", string(mama))

		fmt.Printf("- Creating new account %q\n", prodName)
		out = append(out, newAccount, buyRAMBytes, delegateBW)

		if op.TestnetEnrichProducers {
			fmt.Printf("  DEBUG: Enriching producer %q\n", prodName)
			out = append(out, token.NewTransfer(AN("eosio"), prodName, eos.NewEOSAsset(1000000000), "Hey, make good use of it!"))
		}
	}
	return
}

//

type OpRegisterProducers struct {
}

func (op *OpRegisterProducers) ResetTestnetOptions() {
}

func (op *OpRegisterProducers) Actions(b *BIOS) (out []*eos.Action, err error) {
	for _, prod := range b.Network.OrderedPeers(b.Network.MyNetwork()) {
		prodName := prod.Discovery.TargetAccountName
		if prodName == AN("eosio") {
			prodName = prod.Discovery.SeedNetworkAccountName // only happens with --single
		}

		url := ""
		if len(prod.Discovery.URLs) > 0 {
			url = prod.Discovery.URLs[0]
		}
		regprod := system.NewRegProducer(prodName, prod.Discovery.TargetAppointedBlockProducerSigningKey, url) // overridden just below
		regprod.Authorization[0].Actor = AN("eosio")
		out = append(out, regprod)
	}
	return
}

//

type OpInjectSnapshot struct {
	BuyRAM                  uint64 `json:"buy_ram_bytes"`
	TestnetTruncateSnapshot int    `json:"TESTNET_TRUNCATE_SNAPSHOT"`
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

	fmt.Printf("Preparing %d actions to honor crowdsale holders\n", len(snapshotData))
	for idx, hodler := range snapshotData {
		destAccount := AN(strings.Replace(hodler.AccountName, "0", "genesis", -1)[:12])

		if hodler.EthereumAddress == "0x00000000000000000000000000000000000000b1" {
			// the undelegatebw action does special unvesting for the b1 account
			destAccount = "b1b1b1b1b1b1" // TODO: CONTRACT SHOULD CHANGE TOO
		}

		// fmt.Println("Transfer", hodler, destAccount)

		out = append(out, system.NewNewAccount(AN("eosio"), destAccount, hodler.EOSPublicKey))

		// memo := "Welcome " + hodler.EthereumAddress[len(hodler.EthereumAddress)-6:]

		// out = append(out, token.NewTransfer(AN("eosio"), destAccount, hodler.Balance, memo))

		initialBalance := hodler.Balance // .Sub(eos.NewEOSAsset(int64(op.BuyRAM))) // take ~0.1 to pay for initial RAM
		out = append(out, system.NewBuyRAMBytes(AN("eosio"), destAccount, uint32(op.BuyRAM)))

		firstHalf := initialBalance
		secondHalf := initialBalance

		firstHalf.Amount = firstHalf.Amount / 2
		secondHalf.Amount = hodler.Balance.Amount - firstHalf.Amount

		delBW := system.NewDelegateBW(AN("eosio"), destAccount, firstHalf, secondHalf, true)
		delBW.Authorization[0].Actor = eos.AN("eosio")
		out = append(out, delBW)

		if trunc := op.TestnetTruncateSnapshot; trunc != 0 {
			if idx == trunc {
				fmt.Printf("- DEBUG: truncated snapshot at %d rows\n", trunc)
				break
			}
		}
	}

	return
}

//

type OpSetProds struct{}

func (op *OpSetProds) ResetTestnetOptions() {}
func (op *OpSetProds) Actions(b *BIOS) (out []*eos.Action, err error) {
	// prodkeys := []system.ProducerKey{system.ProducerKey{
	// 	ProducerName:    AN("eosio"),
	// 	BlockSigningKey: b.EphemeralPrivateKey.PublicKey(),
	// }}

	// SHOULD WE `regproducer` here ? or `setprods` is fine ?

	prodkeys := []system.ProducerKey{}
	for _, prod := range b.ShuffledProducers {
		targetKey := prod.Discovery.TargetAppointedBlockProducerSigningKey
		targetAcct := prod.Discovery.TargetAccountName
		if targetAcct == AN("eosio") {
			targetKey = b.EphemeralPublicKey
		}
		prodkeys = append(prodkeys, system.ProducerKey{targetAcct, targetKey})
		if len(prodkeys) >= 21 {
			break
		}
	}
	out = append(out, system.NewSetProds(prodkeys))

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

	nullKey := ecc.PublicKey{Curve: ecc.CurveK1, Content: make([]byte, 33, 33)}
	for _, acct := range op.Accounts {
		out = append(out,
			system.NewUpdateAuth(acct, PN("active"), PN("owner"), eos.Authority{
				Threshold: 1,
				Keys: []eos.KeyWeight{
					{
						PublicKey: nullKey,
						Weight:    1,
					},
				},
			}, PN("active")),
			system.NewUpdateAuth(acct, PN("owner"), PN(""), eos.Authority{
				Threshold: 1,
				Keys: []eos.KeyWeight{
					{
						PublicKey: nullKey,
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
