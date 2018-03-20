package main

import (
	"log"

	"github.com/eosioca/eosapi"
)

type BIOS struct {
	LaunchData        *LaunchData
	Config            *Config
	API               *eosapi.EOSAPI
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
	b.AnnounceAppointedBlockProducers()

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

func (b *BIOS) AnnounceAppointedBlockProducers() {
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
	// Generate keypair
	// Generate genesis.json
	// Produce a config.ini
	// Call webhook ConfigReady
	//   If no such webhook, wait for ENTER keypress after printing the config material.
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
	// Wait on stdin for kickstart data (will we have some other polling / subscription mechanisms?)
	//    Accept any base64, unpadded, multi-line until we receive a blank line, concat and decode.
	// Decrypt the Kickstart data
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
	// Wait on stdin
	//   Input should be simply the p2p endpoint of any node that initialized
	// It'll be an armored GPG-signed (base64) blob containing each producer's `Kickstart Data`, relaying the original `PrivateKeyUsed`, but with their own `p2p_address`
	// Dispatch webhook ConnectToBIOS, relaying the `PrivateKeyUsed` discovered by the ABPs
	// We can then run the same verifications, without sabotage being enabled or risked.
	// At this point, our node is sync'd with the network
	// We call `regproducer` for ourselves, since we want to register don't we ?
	return nil
}

/// Setup

func (b *BIOS) ShuffleProducers(btcMerkleRoot []byte) error {
	// we'll shuffle later :)
	if b.Config.NoShuffle {
		b.ShuffledProducers = b.LaunchData.Producers
	} else {
		// FIXME: put an algorithm here..
		b.ShuffledProducers = b.LaunchData.Producers
	}
	return nil
}

func (b *BIOS) IsBootNode(account string) bool {
	return b.ShuffledProducers[0].EOSIOAccountName == account
}

func (b *BIOS) AmIBootNode() bool {
	return b.IsBootNode(b.Config.Producer.MyAccount)
}

func (b *BIOS) IsAppointedBlockProducer(account string) bool {
	for i := 1; i < 22 && len(b.ShuffledProducers) > i; i++ {
		if b.ShuffledProducers[i].EOSIOAccountName == account {
			return true
		}
	}
	return false
}

func (b *BIOS) AmIAppointedBlockProducer() bool {
	return b.IsAppointedBlockProducer(b.Config.Producer.MyAccount)
}
