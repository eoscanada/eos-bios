// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/eoscanada/eos-bios"
	"github.com/eoscanada/eos-bios/disco"
	"github.com/spf13/cobra"
)

// discoverCmd represents the discovery command
var discoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Discover and update info about all peers in the network, based on an initial discovery URL",
	Long:  `This uses the "network.seed_discovery_url" key in your configuration to start discovery.`,
	Run: func(cmd *cobra.Command, args []string) {

		ipfs, err := bios.NewIPFS(ipfsLocalGatewayAddress, ipfsGatewayAddress)
		if err != nil {
			fmt.Println("ipfs client error:", err)
			os.Exit(1)
		}

		var discovery *disco.Discovery
		if discovery, err = bios.LoadDiscoveryFromFile(myDiscoveryFile); err != nil {
			fmt.Println("error")
			fmt.Fprintf(os.Stderr, "format invalid: %s", err)
			os.Exit(1)
		}

		api, err := api(discovery.SeedNetworkChainID)
		if err != nil {
			fmt.Println("api error:", err)
			os.Exit(1)
		}
		net, err := fetchNetwork(api, ipfs, seedNetworkContract)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Fetched successfully")

		net.PrintOrderedPeers()
	},
}

func init() {
	RootCmd.AddCommand(discoverCmd)
}
