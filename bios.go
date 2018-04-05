package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/eosioca/eosapi"
	"github.com/eosioca/eosapi/ecc"
	"github.com/eosioca/eosapi/system"
	"github.com/eosioca/eosapi/token"
)

type BIOS struct {
	LaunchData   *LaunchData
	Config       *Config
	API          *eos.EOSAPI
	Snapshot     Snapshot
	ShuffleBlock struct {
		Time       time.Time
		MerkleRoot []byte
	}
	ShuffledProducers []*ProducerDef
	MyProducer        *ProducerDef
}

func NewBIOS(launchData *LaunchData, config *Config, snapshotData Snapshot, api *eos.EOSAPI) *BIOS {
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
		return fmt.Errorf("find my producer definition: %s", err)
	}
	b.MyProducer = myProducerDef

	if err := b.DispatchInit(); err != nil {
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

	//b.API.Debug = true

	pubKey := ephemeralPrivateKey.PublicKey().String()
	privKey := ephemeralPrivateKey.String()

	fmt.Println("Generated ephemeral private keys:", pubKey, privKey)

	// Store keys in wallet, to sign `SetCode` and friends..
	if err := b.API.Signer.ImportPrivateKey(privKey); err != nil {
		return fmt.Errorf("ImportWIF: %s", err)
	}

	genesisData := b.GenerateGenesisJSON(pubKey)

	if err = b.DispatchConfigReady(genesisData, "eosio", pubKey, privKey, true); err != nil {
		return fmt.Errorf("dispatch config_ready hook: %s", err)
	}

	for _, prod := range b.ShuffledProducers {
		fmt.Println("Creating new account for", prod.EOSIOAccountName)
		_, err = b.API.SignPushActions(
			system.NewNewAccount(AN("eosio"), prod.EOSIOAccountName, prod.pubKey),
		)
		if err != nil {
			return fmt.Errorf("newaccount %s: %s", prod.EOSIOAccountName, err)
		}
	}

	fmt.Println("Creating account eosio.msig")
	_, err = b.API.SignPushActions(
		system.NewNewAccount(AN("eosio"), AN("eosio.msig"), ephemeralPrivateKey.PublicKey()),
	)
	if err != nil {
		return fmt.Errorf("newaccount eosio.msig: %s", err)
	}

	fmt.Println("Creating account eosio.token")
	_, err = b.API.SignPushActions(
		system.NewNewAccount(AN("eosio"), AN("eosio.token"), ephemeralPrivateKey.PublicKey()),
	)
	if err != nil {
		return fmt.Errorf("newaccount eosio.token: %s", err)
	}

	// Inject bios

	fmt.Println("Setting eosio.bios code for account eosio")
	setCode, err := system.NewSetCodeTx(AN("eosio"), b.Config.BIOSContract.CodePath, b.Config.BIOSContract.ABIPath)
	if err != nil {
		return fmt.Errorf("NewSetCodeTx eosio.bios: %s", err)
	}

	_, err = b.API.SignPushTransaction(setCode, eos.TxOptions{})
	if err != nil {
		return fmt.Errorf("signpushtx code eosio.bios: %s", err)
	}

	// Setpriv on `eosio` and `eosio.msig`

	fmt.Println("Setting privileged account for eosio and eosio.msig")
	_, err = b.API.SignPushActions(
		system.NewSetPriv(AN("eosio")),
		system.NewSetPriv(AN("eosio.msig")),
	)
	if err != nil {
		return fmt.Errorf("setpriv eosio: %s", err)
	}

	// Inject msig code

	fmt.Println("Setting eosio.msig code for account eosio.msig")
	setCode, err = system.NewSetCodeTx(AN("eosio.msig"), b.Config.MsigContract.CodePath, b.Config.MsigContract.ABIPath)
	if err != nil {
		return fmt.Errorf("NewSetCodeTx eosio.msig: %s", err)
	}

	fmt.Println(" - code")
	acts := setCode.Actions[:]
	setCode.Actions = acts[:1]
	_, err = b.API.SignPushTransaction(setCode, eos.TxOptions{})
	if err != nil {
		return fmt.Errorf("signpushtx code eosio.msig: %s", err)
	}
	// FIXME: the abi isn't ready yet.. it doesn't serialize properly, yields something invalid.
	// fmt.Println(" - abi")
	// setCode.Actions = acts[1:]
	// _, err = b.API.SignPushTransaction(setCode, eos.TxOptions{})
	// if err != nil {
	// 	return fmt.Errorf("signpushtx abi eosio.msig: %s", err)
	// }

	// Inject eosio.token code

	fmt.Println("Setting eosio.token code for account eosio.token")
	setCode, err = system.NewSetCodeTx(AN("eosio.token"), b.Config.TokenContract.CodePath, b.Config.TokenContract.ABIPath)
	if err != nil {
		return fmt.Errorf("NewSetCodeTx eosio.token: %s", err)
	}

	_, err = b.API.SignPushTransaction(setCode, eos.TxOptions{})
	if err != nil {
		return fmt.Errorf("signpushtx code eosio.token: %s", err)
	}

	// See tests/chain_tests/bootseq_tests.cpp and friends..

	fmt.Println("Creating the `EOS` currency symbol")
	_, err = b.API.SignPushActions(
		token.NewCreate(AN("eosio"), eos.Asset{Amount: 10000000000000, Symbol: eos.EOSSymbol}, false, false, false),
	)
	if err != nil {
		return fmt.Errorf("create token: %s", err)
	}

	// TODO: Issue from the `eosio.token` contract.. `transfer` and
	// `issue` on `eosio.system` is probably going to disappear.
	fmt.Println("Issuing base currency as EOS")
	_, err = b.API.SignPushActions(
		token.NewIssue(AN("eosio"), eos.Asset{Amount: 10000000000000, Symbol: eos.EOSSymbol}, "Initial issuance"),
	)
	if err != nil {
		return fmt.Errorf("issue: %s", err)
	}

	for idx, hodler := range b.Snapshot {
		destAccount := AN("genesis." + strings.Trim(eos.NameToString(uint64(idx+1)), "."))
		fmt.Println("Transfer", hodler, destAccount)

		_, err = b.API.SignPushActions(
			system.NewNewAccount(AN("eosio"), destAccount, hodler.EOSPublicKey),
		)
		if err != nil {
			return fmt.Errorf("hodler: newaccount: %s", err)
		}

		memo := "Welcome " + hodler.EthereumAddress[len(hodler.EthereumAddress)-6:]
		_, err := b.API.SignPushActions(
			token.NewTransfer(AN("eosio"), destAccount, hodler.Balance, memo),
		)
		if err != nil {
			return fmt.Errorf("hodler: transfer: %s", err)
		}

		if idx == 5 {
			fmt.Println("- Skipping Transfers")
			break
		}
	}

	// Call SetProds, to setup the first producers.
	fmt.Println("Setting the first batch of producers")
	var prodkeys []system.ProducerKey
	for _, prod := range b.ShuffledProducers {
		prodkeys = append(prodkeys, system.ProducerKey{prod.EOSIOAccountName, prod.pubKey})
	}
	_, err = b.API.SignPushActions(system.NewSetProds(0, prodkeys))
	if err != nil {
		return fmt.Errorf("setprods: %s", err)
	}

	fmt.Println("Replacing eosio account from eosio.bios contract to eosio.system")
	_, err = b.API.SetCode(AN("eosio"), b.Config.SystemContract.CodePath, b.Config.SystemContract.ABIPath)
	if err != nil {
		return fmt.Errorf("setcode: %s", err)
	}

	fmt.Println("Disabling authorization for accounts eosio, eosio.msig and eosio.token")
	_, err = b.API.SignPushActions(
		system.NewUpdateAuth(AN("eosio"), PN("active"), PN("owner"), eos.Authority{Threshold: 0}, PN("active")),
		system.NewUpdateAuth(AN("eosio"), PN("owner"), PN(""), eos.Authority{Threshold: 0}, PN("owner")),
		system.NewUpdateAuth(AN("eosio.msig"), PN("active"), PN("owner"), eos.Authority{Threshold: 0}, PN("active")),
		system.NewUpdateAuth(AN("eosio.msig"), PN("owner"), PN(""), eos.Authority{Threshold: 0}, PN("owner")),
		system.NewUpdateAuth(AN("eosio.token"), PN("active"), PN("owner"), eos.Authority{Threshold: 0}, PN("active")),
		system.NewUpdateAuth(AN("eosio.token"), PN("owner"), PN(""), eos.Authority{Threshold: 0}, PN("owner")),
	)
	if err != nil {
		return fmt.Errorf("updateauth: %s", err)
	}

	kickstartData := &KickstartData{
		BIOSP2PAddress: b.Config.Producer.SecretP2PAddress,
		PublicKeyUsed:  pubKey,
		PrivateKeyUsed: privKey,
		GenesisJSON:    genesisData,
	}
	kd, _ := json.Marshal(kickstartData)
	ksdata := base64.RawStdEncoding.EncodeToString(kd)

	fmt.Println("PUBLISH THIS KICKSTART DATA:", string(ksdata))

	return b.DispatchDone()
	// Create the `Kickstart data`
	// Call webhook PublishKickstartEncrypted
	//   or display it on screen for it to be manually disseminated
	// Call `regproducer` for myself now
	// Return and we're done.
	// Dispatch WebhookBIOSNodeDone
}

func (b *BIOS) RunABPStage1() error {
	fmt.Println("Waiting on kickstart data from the BIOS Node. Check their social presence!")

	// Wait on stdin for kickstart data (will we have some other polling / subscription mechanisms?)
	//    Accept any base64, unpadded, multi-line until we receive a blank line, concat and decode.
	// FIXME: this is a quick hack to just pass the p2p address
	lines, err := ScanLinesUntilBlank()
	if err != nil {
		return err
	}

	rawKickstartData, err := base64.RawStdEncoding.DecodeString(strings.Replace(strings.TrimSpace(lines), "\n", "", -1))
	if err != nil {
		return fmt.Errorf("kickstrat base64 decode: %s", err)
	}

	var kickstart KickstartData
	err = json.Unmarshal(rawKickstartData, &kickstart)
	if err != nil {
		return fmt.Errorf("unmarshal kickstart data: %s", err)
	}

	// Decrypt the Kickstart data
	//   Do extensive validation on the input (tight regexp for address, for private key?)

	err = b.DispatchConnectToBIOS(kickstart, func() error {
		_, err := b.API.NetConnect(kickstart.BIOSP2PAddress)
		return err
	})
	if err != nil {
		return err
	}

	for {
		time.Sleep(1 * time.Second)
		acct, err := b.API.GetAccount(AN("eosio"))
		if err != nil {
			fmt.Printf("e")
			continue
		}

		if len(acct.Permissions) != 2 || acct.Permissions[0].RequiredAuth.Threshold != 0 || acct.Permissions[1].RequiredAuth.Threshold != 0 {
			// FIXME: perhaps check that there are no keys and
			// accounts.. that the account is *really* disabled.  we
			// can check elsewhere though.
			fmt.Printf(".")
			continue
		}

		fmt.Printf(" good! BIOS signed off, chain is sync'd!")
		break
	}

	_, err = b.API.SignPushActions(system.NewRegProducer(b.MyProducer.EOSIOAccountName, b.MyProducer.pubKey, b.Config.MyParameters))
	if err != nil {
		return fmt.Errorf("setprods: %s", err)
	}

	// Do all the checks:
	//  - all Producers are properly setup
	//  - anything fails, SABOTAGE
	// We call `regproducer` for ourselves.
	// Publish a PGP Signed message with your local IP.. push to properties
	// Dispatch webhook PublishKickstartPublic (with a Kickstart Data object)

	return b.DispatchDone()
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
