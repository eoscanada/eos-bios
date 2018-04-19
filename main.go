package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eoscanada/eos-go"
)

var localConfig = flag.String("local-config", "", "Local .yaml configuration file.")
var launchData = flag.String("launch-data", "launch.yaml", "Path to a launch.yaml file, your community-agreed ignition configuration.")
var versionFlag = flag.Bool("version", false, "Show the version and quit. Hint hint, it's: "+version)
var version string

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println("eos-bios version:", version)
		os.Exit(0)
	}

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

	api := eos.New(config.Producer.apiAddressURL, chainID)
	if err != nil {
		log.Fatalln("producer node error:", err)
	}

	api.SetSigner(eos.NewKeyBag())

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

	if err = bios.setMyProducerDefs(); err != nil {
		log.Fatalln("Failed to get my producer definition:", err)
	}

	if err := bios.Run(); err != nil {
		log.Fatalf("ERROR RUNNING BIOS: %s", err)
	}

	fmt.Printf("Done at UTC %s\n", time.Now().UTC())
}
