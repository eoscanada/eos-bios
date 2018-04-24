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
)

// joinCmd represents the join command
var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Triggers the hooks to join an already running network",
	Long:  `This will run the "join_network" hook with data discovered from the network pointed to by the seed_discovery_url.`,
	Run: func(cmd *cobra.Command, args []string) {
		net, err := fetchNetwork(false)
		if err != nil {
			fmt.Println("error loading network:", err)
			os.Exit(1)
		}

		// TODO: if the network doesn't contain the `genesis.json` at all, ask for the kickstart data on the command line, and use that to boot.

		b := bios.NewBIOS(net, biosConfig, nil)
		if err := b.Init(); err != nil {
			log.Fatalf("BIOS initialization error: %s", err)
		}

		// TODO: only run the `boot` part, ensure we're the BIOS Boot Node..
		if err := b.Run(bios.RoleParticipant); err != nil {
			log.Fatalf("ERROR RUNNING BIOS: %s", err)
		}

		// Fetch the kickstart data through the discovery_url (the genesis should be present somewhere on the network)
		// run b.DispatchConnectAsParticipant(kickstart, mypeer)
		//
		fmt.Println("join called")
	},
}

func init() {
	RootCmd.AddCommand(joinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
