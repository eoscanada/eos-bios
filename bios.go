package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/eosioca/eosapi"
	"github.com/eosioca/eosapi/ecc"
)

type BIOS struct {
	LaunchData   *LaunchData
	Config       *Config
	API          *eosapi.EOSAPI
	Snapshot     Snapshot
	ShuffleBlock struct {
		Time       time.Time
		MerkleRoot []byte
	}
	ShuffledProducers []*ProducerDef
}

func NewBIOS(launchData *LaunchData, config *Config, snapshotData Snapshot, api *eosapi.EOSAPI) *BIOS {
	b := &BIOS{
		LaunchData: launchData,
		Config:     config,
		API:        api,
		Snapshot:   snapshotData,
	}
	return b
}

func (b *BIOS) Run() error {
	fmt.Println("Start BIOS process", time.Now())

	myProducerDef, err := b.MyProducerDef()
	if err != nil {
		return err
	}

	// TODO: We need to make SURE we *DO* have signing keys for the myProducerDef.EOSIOPublicKey,
	// by checking the linked wallet or som'thn..

	if err := b.DispatchInit(b.GenerateGenesisJSON(myProducerDef.EOSIOPublicKey)); err != nil {
		return fmt.Errorf("failed init hook: %s", err)
	}

	// Main program entrypoint, called when setup is done.
	b.PrintAppointedBlockProducers()

	if b.AmIBootNode() {
		if err := b.RunBootNodeStage1(); err != nil {
			return fmt.Errorf("boot node stage1: %s", err)
		}
	} else if b.AmIAppointedBlockProducer() {
		if err := b.RunABPStage1(); err != nil {
			return fmt.Errorf("abp stage1: %s", err)
		}
	} else {
		if err := b.WaitStage1End(); err != nil {
			return fmt.Errorf("waiting stage1: %s", err)
		}
	}

	fmt.Println("BIOS Run done")

	return nil
}

func (b *BIOS) PrintAppointedBlockProducers() {
	if b.AmIBootNode() {
		fmt.Println("STAGE 0: I AM THE BOOT NODE! Let's get the ball rolling.")

	} else if b.AmIAppointedBlockProducer() {
		fmt.Println("STAGE 0: I am NOT the BOOT NODE, but I AM ONE of the Appointed Block Producers. Stay tuned and watch the boot node's media properties.")
	} else {
		fmt.Println("STAGE 0: hrm.. I'm not part of the Appointed Block Producers, let's wait and be ready to join")
	}

	fmt.Printf("BIOS NODE: %s\n", b.ShuffledProducers[0].String())
	for i := 1; i < 22 && len(b.ShuffledProducers) > i; i++ {
		fmt.Printf("ABP %02d:    %s\n", i, b.ShuffledProducers[i].String())
	}
}

func (b *BIOS) RunBootNodeStage1() error {
	ephemeralPrivateKey, err := b.GenerateEphemeralPrivKey()
	if err != nil {
		return err
	}

	pubKey := ephemeralPrivateKey.PublicKey().String()
	privKey := ephemeralPrivateKey.String()

	genesisData := b.GenerateGenesisJSON(pubKey)

	fmt.Println("Triggering `config_ready` hook")

	if err = b.DispatchConfigReady(genesisData, "eosio", pubKey, privKey, true); err != nil {
		return fmt.Errorf("dispatch config_ready hook: %s", err)
	}

	eosioAcct := eosapi.AccountName("eosio")
	_, err = b.API.SetCode(eosioAcct, b.Config.SystemContract.CodePath, b.Config.SystemContract.ABIPath)
	if err != nil {
		return fmt.Errorf("setcode: %s", err)
	}

	for _, prod := range b.ShuffledProducers {
		_, err = b.API.NewAccount(eosioAcct, prod.EOSIOAccountName, prod.pubKey)
		if err != nil {
			return fmt.Errorf("newaccount %s: %s", prod.EOSIOAccountName, err)
		}
	}

	eosSymbol := eosapi.Symbol{Precision: 4, Symbol: "EOS"}

	_, err = b.API.Issue("eosio", eosapi.Asset{Amount: 10000000000000, Symbol: eosSymbol})
	if err != nil {
		return fmt.Errorf("issue: %s", err)
	}

	for idx, hodler := range b.Snapshot {
		// 7108558431253954560 stringifies to `genesis.`
		destAccount := eosapi.AccountName(eosapi.NameToString(uint64(7108558431253954560) + uint64(idx)))
		fmt.Println("Transfer", hodler, destAccount)

		qty := eosapi.Asset{Amount: 10000000000000, Symbol: eosSymbol}
		_, err := b.API.Transfer(eosapi.AccountName("eosio"), destAccount, qty, "Welcome "+hodler.EthereumAddress[len(hodler.EthereumAddress)-6:])
		if err != nil {
			return fmt.Errorf("transfer: %s", err)
		}
	}

	// Call `transfer` for everyone in `snapshot.csv`
	// Call `updateauth` and trash the `eosio` account.
	// Create the `Kickstart data`
	// Call webhook PublishKickstartEncrypted
	//   or display it on screen for it to be manually disseminated
	// Call `regproducer` for myself now
	// Return and we're done.
	// Dispatch WebhookBIOSNodeDone
	return nil
}

func (b *BIOS) RunABPStage1() error {
	fmt.Println("Waiting on kickstart data from the BIOS Node. Check their social presence!")

	// Wait on stdin for kickstart data (will we have some other polling / subscription mechanisms?)
	//    Accept any base64, unpadded, multi-line until we receive a blank line, concat and decode.
	// Decrypt the Kickstart data
	//   Do extensive validation on the input (tight regexp for address, for private key?)
	// Call `api.NetConnect()` on the `p2p_address` therein.
	// Dispatch Webhook ConnectToBIOS
	//   Display `config.ini` snippets to inject and wait on keypress.
	// Poll your P2P-Address, until the network syncs..
	// Do all the checks:
	//  - all Producers are properly setup
	//  - anything fails, SABOTAGE
	// We call `regproducer` for ourselves.
	// Publish a PGP Signed message with your local IP.. push to properties
	// Dispatch webhook PublishKickstartPublic (with a Kickstart Data object)
	return nil
}

func (b *BIOS) WaitStage1End() error {
	fmt.Println("Waiting for Appointed Block Producers to finish their jobs. Check their social presence!")
	// Wait on stdin
	//   Input should be simply the p2p endpoint of any node that initialized
	// It'll be an armored GPG-signed (base64) blob containing each producer's `Kickstart Data`, relaying the original `PrivateKeyUsed`, but with their own `p2p_address`
	//   Again, do extensive validation on the input, anything reaching webhooks.
	// Dispatch webhook ConnectToBIOS, relaying the `PrivateKeyUsed` discovered by the ABPs
	// We can then run the same verifications, without sabotage being enabled or risked.
	// At this point, our node is sync'd with the network
	// We call `regproducer` for ourselves, since we want to register don't we ?
	return nil
}

func (b *BIOS) GenerateEphemeralPrivKey() (*ecc.PrivateKey, error) {
	return ecc.NewRandomPrivateKey()
}

func (b *BIOS) GenerateGenesisJSON(pubKey string) string {
	cnt, _ := json.Marshal(&GenesisJSON{
		InitialTimestamp: b.ShuffleBlock.Time.UTC().Format("2006-01-02T15:04:05"),
		InitialKey:       pubKey,
		InitialChainID:   hex.EncodeToString(b.API.ChainID),
	}) // known not to fail
	return string(cnt)
}

/// Setup

func (b *BIOS) ShuffleProducers(btcMerkleRoot []byte, blockTime time.Time) error {
	// we'll shuffle later :)
	if b.Config.NoShuffle {
		b.ShuffledProducers = b.LaunchData.Producers
		b.ShuffleBlock.Time = time.Now().UTC()
		b.ShuffleBlock.MerkleRoot = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	} else {
		// FIXME: put an algorithm here..
		b.ShuffledProducers = b.LaunchData.Producers
		b.ShuffleBlock.Time = blockTime
		b.ShuffleBlock.MerkleRoot = btcMerkleRoot
	}
	return nil
}

func (b *BIOS) IsBootNode(account string) bool {
	return string(b.ShuffledProducers[0].EOSIOAccountName) == account
}

func (b *BIOS) AmIBootNode() bool {
	return b.IsBootNode(b.Config.Producer.MyAccount)
}

func (b *BIOS) IsAppointedBlockProducer(account string) bool {
	for i := 1; i < 22 && len(b.ShuffledProducers) > i; i++ {
		if string(b.ShuffledProducers[i].EOSIOAccountName) == account {
			return true
		}
	}
	return false
}

func (b *BIOS) AmIAppointedBlockProducer() bool {
	return b.IsAppointedBlockProducer(b.Config.Producer.MyAccount)
}

func (b *BIOS) MyProducerDef() (*ProducerDef, error) {
	for _, prod := range b.LaunchData.Producers {
		if b.Config.Producer.MyAccount == string(prod.EOSIOAccountName) {
			return prod, nil
		}
	}
	return nil, fmt.Errorf("no config found")
}
