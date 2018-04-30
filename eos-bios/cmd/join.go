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
	"log"
	"os"

	bios "github.com/eoscanada/eos-bios"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var verifyFlag bool

// joinCmd represents the join command
var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Triggers the hooks to join an already running network",
	Long:  `This will run the "join_network" hook with data discovered from the network pointed to by the seed_discovery_url.`,
	Run: func(cmd *cobra.Command, args []string) {
		ipfs, err := bios.NewIPFS(ipfsGatewayAddress, ipfsLocalGatewayAddress)
		if err != nil {
			fmt.Println("ipfs client error:", err)
			os.Exit(1)
		}

		net, err := fetchNetwork(ipfs)
		if err != nil {
			fmt.Println("error loading network:", err)
			os.Exit(1)
		}

		b := bios.NewBIOS(net, nil)
		if err := b.Init(); err != nil {
			log.Fatalf("BIOS initialization error: %s", err)
		}

		if err := b.StartJoin(verifyFlag); err != nil {
			log.Fatalf("error joining network: %s", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(joinCmd)

	joinCmd.Flags().BoolVarP(&verifyFlag, "verify", "v", false, "Verify the boot sequence by comparing all expected actions against what is on the first blocks of the chain")
	joinCmd.Flags().StringVarP(&apiAddress, "api-address", "", "http://localhost:8888", "RPC endpoint of your nodeos instance. Needs only to be reachable by this process.")

	for _, flag := range []string{"verify", "api-address"} {
		viper.BindPFlag(flag, joinCmd.Flags().Lookup(flag))
	}
}
