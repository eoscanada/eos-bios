package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/eosioca/eosapi"
)

var localConfig = flag.String("local-config", "", "Local .yaml configuration file.")
var launchData = flag.String("launch-data", "launch.yaml", "Path to a launch.yaml file, your community-agreed ignition configuration.")

func main() {
	flag.Parse()

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

	//
	chainID := "0000000000000000000000000000000000000000000000000000000000000000"

	// chainID will become the HASH of the Constitution, we could start with a sample constitution and hash it ? waddayouthink ?
	api, err := eosapi.New(config.Producer.APIAddress, chainID)
	if err != nil {
		log.Fatalln("producer node error:", err)
	}

	wallet, err := eosapi.New(config.Producer.WalletAddress, chainID)
	if err != nil {
		log.Fatalln("wallet api:", err)
	}

	api.SetSigner(eosapi.NewWalletSigner(wallet))

	// Checking producer node
	info, err := api.GetInfo()
	if err != nil {
		log.Fatalf("Producer node not accessible: %s", err)
	}

	log.Println("Server version:", info.ServerVersion)
	if info.HeadBlockNum > 0 {
		log.Fatalf("Your node is at block %d, aren't we booting a new network?", info.HeadBlockNum)
		os.Exit(1)
	}

	// Checking wallet node
	info, err = wallet.GetInfo()
	if err != nil {
		log.Fatalf("Wallet node not accessible: %s", err)
	}

	// Start the process

	fmt.Println("More things to come...")
}
