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
	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
	"github.com/eoscanada/eos-go/system"
	"github.com/spf13/cobra"
)

// inviteCmd represents the invite command
var inviteCmd = &cobra.Command{
	Use:   "invite [account_name] [public_key]",
	Short: "Invite a fellow block producer to the seed network where you have access to",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Reading discovery file... ")
		var discovery *disco.Discovery
		var err error
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
		publicKey, err := ecc.NewPublicKey(args[1])
		if err != nil {
			fmt.Println("error")
			fmt.Fprintf(os.Stderr, "public key: %s", err)
			os.Exit(1)
		}
		api.SignPushActions(
			system.NewNewAccount(
				eos.AccountName(discovery.SeedNetworkAccountName),
				eos.AccountName(args[0]),
				publicKey,
			),
		)
		fmt.Println("inviting", args[0], "with public key", args[1])
		fmt.Println("done ... chicken")
	},
}

func init() {
	RootCmd.AddCommand(inviteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inviteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inviteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
