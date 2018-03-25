package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/eosioca/eosapi"
	"github.com/eosioca/eosapi/ecc"
)

type BIOS struct {
	LaunchData   *LaunchData
	Config       *Config
	API          *eosapi.EOSAPI
	ShuffleBlock struct {
		Time       time.Time
		MerkleRoot []byte
	}
	ShuffledProducers []*ProducerDef
}

func NewBIOS(launchData *LaunchData, config *Config, api *eosapi.EOSAPI) *BIOS {
	b := &BIOS{
		LaunchData: launchData,
		Config:     config,
		API:        api,
	}
	return b
}

func (b *BIOS) Run() error {
	// Main program entrypoint, called when setup is done.
	b.PrintAppointedBlockProducers()

	if b.AmIBootNode() {
		if err := b.RunBootNodeStage1(); err != nil {
			return err
		}
	} else if b.AmIAppointedBlockProducer() {
		if err := b.RunABPStage1(); err != nil {
			return err
		}
	} else {
		if err := b.WaitStage1End(); err != nil {
			return err
		}
	}

	return nil
}

func (b *BIOS) PrintAppointedBlockProducers() {
	if b.AmIBootNode() {
		log.Println("STAGE 0: I AM THE BOOT NODE! Let's get the ball rolling.")

	} else if b.AmIAppointedBlockProducer() {
		log.Println("STAGE 0: I am NOT the BOOT NODE, but I AM ONE of the Appointed Block Producers. Stay tuned and watch the boot node's media properties.")
	} else {
		log.Println("STAGE 0: hrm.. I'm not part of the Appointed Block Producers, let's wait and be ready to join")
	}

	log.Printf("BIOS NODE: %s\n", b.ShuffledProducers[0].String())
	for i := 1; i < 22 && len(b.ShuffledProducers) > i; i++ {
		log.Printf("ABP %02d:  %s\n", i, b.ShuffledProducers[i].String())
	}
}

func (b *BIOS) RunBootNodeStage1() error {
	ephemeralPrivateKey, err := b.GenerateEphemeralPrivKey()
	if err != nil {
		return err
	}

	pubKey := ephemeralPrivateKey.PublicKey().String()
	privKey := ephemeralPrivateKey.String()

	genesisData := b.GenerateGenesisJSON(privKey)

	log.Println("This is your genesis data:", genesisData)

	log.Println("Triggering `config_ready` hook")

	err = b.DispatchConfigReady(genesisData, pubKey, privKey)
	if err != nil {
		return err
	}

	log.Println("If working manually (with no hooks), add these bits to your `config.ini`:")
	fmt.Printf("\n")
	fmt.Printf("producer-name = %q\n", b.ShuffledProducers[0].EOSIOAccountName)
	fmt.Printf("private-key = [%q, %q]\n", pubKey, privKey)
	fmt.Printf("enable-stale-production = true\n")
	fmt.Printf("plugin = eosio::producer_plugin\n")
	fmt.Printf("\n")

	eosioAcct := eosapi.AccountName("eosio")
	_, err = b.API.SetCode(eosioAcct, b.Config.SystemContract.CodePath, b.Config.SystemContract.ABIPath)
	if err != nil {
		return err
	}

	for _, prod := range b.ShuffledProducers {
		_, err = b.API.NewAccount(eosioAcct, prod.EOSIOAccountName, prod.pubKey)
		if err != nil {
			return err
		}
	}
	// TODO:   If no such webhook, wait for ENTER keypress after printing the config material.
	// Call `setcode` and inject system contract
	// Call `newaccount` for all producers listed in b.LaunchData
	// Call `issue` for everyone in `snapshot.csv`
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
	log.Println("Waiting on kickstart data from the BIOS Node. Check their social presence!")

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
	log.Println("Waiting for Appointed Block Producers to finish their jobs. Check their social presence!")
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

func (b *BIOS) GenerateGenesisJSON(privKey string) string {
	cnt, _ := json.Marshal(&GenesisJSON{
		InitialTimestamp: b.ShuffleBlock.Time.UTC().Format("2006-01-02T15:04:05"),
		InitialKey:       privKey,
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
