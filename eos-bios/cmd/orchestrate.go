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
	"net/url"
	"os"

	bios "github.com/eoscanada/eos-bios"
	eos "github.com/eoscanada/eos-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// orchestrateCmd represents the orchestrate command
var orchestrateCmd = &cobra.Command{
	Use:   "orchestrate",
	Short: "Automate all the operations to launch a new network, by collaborating with other in the launch.",
	Long:  `This operation will auto-select the roles, based on a discovered Network shared amongst participants.`,
	Run: func(cmd *cobra.Command, args []string) {
		ipfs, err := bios.NewIPFS(ipfsGatewayAddress, ipfsLocalGatewayAddress)
		if err != nil {
			fmt.Println("ipfs client error:", err)
			os.Exit(1)
		}

		net, err := fetchNetwork(ipfs)
		if err != nil {
			log.Fatalln("fetch network:", err)
		}

		apiAddressURL, err = url.Parse(apiAddress)
		if err != nil {
			log.Fatalln("error parsing --api-address:", err)
		}

		api := eos.New(apiAddressURL, net.ChainID())
		api.SetSigner(eos.NewKeyBag())

		b := bios.NewBIOS(net, api)

		if err := b.Init(); err != nil {
			log.Fatalf("BIOS initialization error: %s", err)
		}

		if err := b.StartOrchestrate(secretP2PAddress); err != nil {
			log.Fatalf("error orchestrating: %s", err)
		}

	},
}

func init() {
	RootCmd.AddCommand(orchestrateCmd)

	orchestrateCmd.Flags().StringVarP(&secretP2PAddress, "secret-p2p-address", "", "localhost:9876", "Address to publish once boot is complete. In an orchestrated boot, you would want to keep this one secret to avoid being DDoS'd.")
	orchestrateCmd.Flags().StringVarP(&apiAddress, "api-address", "", "http://localhost:8888", "RPC endpoint of your nodeos instance. Needs only to be reachable by this process.")

	for _, flag := range []string{"secret-p2p-address", "api-address"} {
		viper.BindPFlag(flag, orchestrateCmd.Flags().Lookup(flag))
	}
}
