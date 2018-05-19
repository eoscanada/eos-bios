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

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
	"github.com/eoscanada/eos-go/system"
	"github.com/spf13/cobra"
)

// inviteCmd represents the invite command
var inviteCmd = &cobra.Command{
	Use:   "invite 'account_name' 'EOSpublickey'",
	Short: "Invite a fellow block producer to the seed network where you have access to",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		net, err := fetchNetwork(false, false)
		if err != nil {
			log.Fatalln("fetch network:", err)
		}

		publicKey, err := ecc.NewPublicKey(args[1])
		if err != nil {
			fmt.Println("invalid public key:", err)
			os.Exit(1)
		}

		fmt.Println("")
		fmt.Printf("Creating account %q with public key %q, using my account %q\n", args[0], args[1], net.MyPeer.Discovery.SeedNetworkAccountName)
		fmt.Println("")

		_, err = net.SeedNetAPI.SignPushActions(
			system.NewNewAccount(
				eos.AccountName(net.MyPeer.Discovery.SeedNetworkAccountName),
				eos.AccountName(args[0]),
				publicKey,
			),
			system.NewBuyRAMBytes(
				eos.AccountName(net.MyPeer.Discovery.SeedNetworkAccountName),
				eos.AccountName(args[0]),
				8192),
			system.NewDelegateBW(
				eos.AccountName(net.MyPeer.Discovery.SeedNetworkAccountName),
				eos.AccountName(args[0]),
				eos.NewEOSAsset(1000),
				eos.NewEOSAsset(1000),
				false,
			),
			/* "actions": [
      {
        "account": "eosio",
        "name": "newaccount",
        "authorization": [
          {
            "actor": "eoscanadacom",
            "permission": "active"
          }
        ],
        "data": "202932c94c833055000000d3757730550100000001000264d39e1bb1fc7f2519046f1c35329f49e21c832eb3fd00a0b1e3433b2323d8b7010000000100000001000264d39e1bb1fc7f2519046f1c35329f49e21c832eb3fd00a0b1e3433b2323d8b701000000"
      },
      {
        "account": "eosio",
        "name": "buyrambytes",
        "authorization": [
          {
            "actor": "eoscanadacom",
            "permission": "active"
          }
        ],
        "data": "202932c94c833055000000d37577305500200000"
      },
      {
        "account": "eosio",
        "name": "delegatebw",
        "authorization": [
          {
            "actor": "eoscanadacom",
            "permission": "active"
          }
        ],
        "data": "202932c94c833055000000d375773055102700000000000004454f5300000000102700000000000004454f530000000001"
      }
    ],
 */
		)
		if err != nil {
			log.Fatalln("creating account:", err)
		}

		fmt.Println("Done.")
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
