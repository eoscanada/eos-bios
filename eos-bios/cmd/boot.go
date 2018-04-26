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
	"log"

	bios "github.com/eoscanada/eos-bios"
	eos "github.com/eoscanada/eos-go"
	"github.com/spf13/cobra"
)

// bootCmd represents the boot command
var bootCmd = &cobra.Command{
	Use:   "boot",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		net, err := fetchNetwork()
		if err != nil {
			log.Fatalln("fetch network:", err)
		}

		if biosConfig.Peer.APIAddressURL == nil {
			log.Fatalln("peer.api_address not found")
		}

		api := eos.New(biosConfig.Peer.APIAddressURL, net.ChainID())
		api.SetSigner(eos.NewKeyBag())

		b := bios.NewBIOS(net, biosConfig, api)

		if err := b.Init(); err != nil {
			log.Fatalf("BIOS initialization error: %s", err)
		}

		if err := b.StartBoot(); err != nil {
			log.Fatalf("error booting network: %s", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(bootCmd)
}
