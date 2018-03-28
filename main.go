package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/eosioca/eosapi"
)

var localConfig = flag.String("local-config", "", "Local .yaml configuration file.")
var launchData = flag.String("launch-data", "launch.yaml", "Path to a launch.yaml file, your community-agreed ignition configuration.")

func main() {
	flag.Parse()

	if *localConfig == "" || *launchData == "" {
		log.Fatalln("missing --launch-data or --local-config")
	}

	config, err := LoadLocalConfig(*localConfig)
	if err != nil {
		log.Fatalln("local config load error:", err)
	}

	launch, err := loadLaunchFile(*launchData, config)
	if err != nil {
		log.Fatalln("launch data error:", err)
	}

	_ = launch.LaunchBitcoinBlockHeight
	// Implement the Bitcoin block fetcher, and merkle root checker..
	//    Implement 3 sources, connect to BTC node, use one of the block explorers, check their APIs.
	// Seed `rand.Seed`

	// chainID will become the HASH of the Constitution? We could
	// start with a sample constitution and hash it ? waddayouthink ?
	chainID := make([]byte, 32, 32)

	api := eosapi.New(config.Producer.apiAddressURL, chainID)
	if err != nil {
		log.Fatalln("producer node error:", err)
	}

	wallet := eosapi.New(config.Producer.walletAddressURL, chainID)
	if err != nil {
		log.Fatalln("wallet api:", err)
	}

	// FIXME: when ECC signatures work natively in Go, we can use the
	// `eosapi.KeyBag` signer instead.
	api.SetSigner(eosapi.NewWalletSigner(wallet))

	// Checking wallet node

	_, err = wallet.WalletPublicKeys()
	if err != nil {
		log.Fatalf("Wallet node not accessible: %s", err)
	}

	// Load the snapshot.csv
	snapshotData, err := NewSnapshot(config.OpeningBalances.SnapshotPath)
	if err != nil {
		log.Fatalln("Failed loading snapshot csv:", err)
	}

	// Start BIOS
	bios := NewBIOS(launch, config, snapshotData, api)

	// FIXME: replace by the BTC data.
	err = bios.ShuffleProducers([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, time.Now().UTC())
	if err != nil {
		log.Fatalln("Failed shuffling:", err)
	}

	if err := bios.Run(); err != nil {
		log.Fatalf("ERROR RUNNING BIOS: %s", err)
	}

	fmt.Printf("Done at %s (UTC %s)\n", time.Now(), time.Now().UTC())
}
