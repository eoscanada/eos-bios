package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/eosioca/eosapi"
)

var producerAPIAddress = flag.String("bp-api-address", "", "Target API endpoint for the locally booting node, a clean-slate node. It can be routable only from the local machine.")
var producerP2PAddress = flag.String("bp-p2p-address", "", "Endpoint which will be published at the end of the process. Needs to be externally routable.")
var eosioMyAccount = flag.String("eosio-my-account", "", "Endpoint which will be published at the end of the process. Needs to be externally routable.")
var eosioSystemCodePath = flag.String("eosio-system-code", "./eosio-system.wasm", "Path to a compiled eosio.system contract .wasm file")
var eosioSystemABIPath = flag.String("eosio-system-abi", "./eosio-system.abi", "Path to an eosio.system .abi file")
var openingBalancesSnapshotPath = flag.String("opening-balances-snapshot", "./snapshot.csv", "Path to a fresh snapshot of the ERC-20 Crowdsale token")
var keybaseKeyPath = flag.String("keybase-key", "", "Path to a PGP key, or keybase thing.. TBD")
var launchData = flag.String("launch-data", "launch.yaml", "Path to the launch.yaml file")

func main() {
	flag.Parse()

	config, err := loadLaunchFile(*launchData, *openingBalancesSnapshotPath, *eosioSystemCodePath, *eosioSystemABIPath)
	if err != nil {
		log.Fatalln("launch data error:", err)
	}

	_ = config.LaunchBitcoinBlockHeight
	// Implement the Bitcoin block fetcher, and merkle root checker..
	//    Implement 3 sources, connect to BTC node, use one of the block explorers, check their APIs.
	// Seed `rand.Seed`

	//

	api := eosapi.New(*producerAPIAddress)
	info, err := api.GetInfo()
	if err != nil {
		log.Fatalf("Local node not accessible: %s", err)
	}

	log.Println("Server version:", info.ServerVersion)
	if info.HeadBlockNum > 0 {
		log.Fatalf("Your node is at block %d, aren't we booting a new network?", info.HeadBlockNum)
		os.Exit(1)
	}

	fmt.Println("More things to come...")
}
