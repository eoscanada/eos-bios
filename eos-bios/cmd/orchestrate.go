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

// orchestrateCmd represents the orchestrate command
var orchestrateCmd = &cobra.Command{
	Use:   "orchestrate",
	Short: "Automate all the operations to launch a new network, by collaborating with other in the launch.",
	Long:  `This operation will auto-select the roles, based on a discovered Network shared amongst participants.`,
	Run: func(cmd *cobra.Command, args []string) {
		net, err := fetchNetwork(false)
		if err != nil {
			log.Fatalln("fetch network:", err)
		}

		if biosConfig.Peer.APIAddressURL == nil {
			log.Fatalln("peer.api_address not found")
		}

		api := eos.New(biosConfig.Peer.APIAddressURL, net.ChainID())
		api.SetSigner(eos.NewKeyBag())

		// Start BIOS
		bios := bios.NewBIOS(net, biosConfig, api)

		if err := bios.Init(); err != nil {
			log.Fatalf("BIOS initialization error: %s", err)
		}

		if err := bios.StartOrchestrate(); err != nil {
			log.Fatalf("error orchestrating: %s", err)
		}

	},
}

func init() {
	RootCmd.AddCommand(orchestrateCmd)
}
