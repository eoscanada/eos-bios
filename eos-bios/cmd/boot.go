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

// bootCmd represents the boot command
var bootCmd = &cobra.Command{
	Use:   "boot",
	Short: "Triggers hooks to boot a new network or node",
	Long: `This will run the "boot_network" hook with data generated locally for a new network.

The "publish_kickstart_data" will also be run, giving you the opportunity to disseminate what is required for people to join your network.

Boot is what happens when you run "eos-bios orchestrate" and you are selected to be the BIOS Boot node.
`,
	Run: func(cmd *cobra.Command, args []string) {
		net, err := fetchNetwork()
		if err != nil {
			log.Fatalln("fetch network:", err)
		}

		ipfs, err := bios.NewIPFS(ipfsAPIAddress, ipfsGatewayAddress, ipfsLocalGatewayAddress)
		if err != nil {
			fmt.Println("ipfs client error:", err)
			os.Exit(1)
		}

		apiAddressURL, err = url.Parse(apiAddress)
		if err != nil {
			log.Fatalln("error parsing --api-address:", err)
		}

		api := eos.New(apiAddressURL, net.ChainID())
		api.SetSigner(eos.NewKeyBag())

		b := bios.NewBIOS(net, api, ipfs)

		if err := b.Init(); err != nil {
			log.Fatalf("BIOS initialization error: %s", err)
		}

		if err := b.StartBoot(secretP2PAddress); err != nil {
			log.Fatalf("error booting network: %s", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(bootCmd)

	bootCmd.Flags().StringVarP(&secretP2PAddress, "secret-p2p-address", "", "localhost:9876", "Address to publish once boot is complete. In an orchestrated boot, you would want to keep this one secret to avoid being DDoS'd.")
	bootCmd.Flags().StringVarP(&apiAddress, "api-address", "", "http://localhost:8888", "RPC endpoint of your nodeos instance. Needs only to be reachable by this process.")

	for _, flag := range []string{"secret-p2p-address", "api-address"} {
		viper.BindPFlag(flag, bootCmd.Flags().Lookup(flag))
	}
}
