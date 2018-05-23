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

	"github.com/eoscanada/eos-bios/bios"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// discoverCmd represents the discovery command
var discoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Discover fetches the network you are participating in on the `eosio.disco` contract.  It does not show networks you are not participating in. Use `list` for that.",
	Run: func(cmd *cobra.Command, args []string) {
		net, err := fetchNetwork(false, false)
		if err != nil {
			log.Fatalln("fetch network:", err)
		}

		if elect := viper.GetString("elect"); elect != "" {
			net.CalculateNetworkWeights(elect)
		}

		net.PrintOrderedPeers()

		if viper.GetBool("serve") {
			bios.Serve(net)
		}
	},
}

func init() {
	RootCmd.AddCommand(discoverCmd)
	discoverCmd.Flags().BoolP("serve", "", false, "Serve the discovery visualization on http://localhost:10101")

	if err := viper.BindPFlag("serve", discoverCmd.Flags().Lookup("serve")); err != nil {
		panic(err)
	}
}
