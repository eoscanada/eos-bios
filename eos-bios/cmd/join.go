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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// joinCmd represents the join command
var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Triggers the hooks to join an already running network",
	Long:  `This will run the "join_network" hook with data discovered from the network pointed to by the seed_discovery_url.`,
	Run: func(cmd *cobra.Command, args []string) {
		net, err := fetchNetwork()
		if err != nil {
			log.Fatalln("fetch network:", err)
		}

		b, err := setupBIOS(net)
		if err != nil {
			log.Fatalln("bios setup:", err)
		}

		if err := b.Init(); err != nil {
			log.Fatalf("BIOS initialization error: %s", err)
		}

		if err := b.StartJoin(viper.GetBool("verify")); err != nil {
			log.Fatalf("error joining network: %s", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(joinCmd)

	joinCmd.Flags().BoolP("verify", "v", false, "Verify the boot sequence by comparing all expected actions against what is on the first blocks of the chain")

	if err := viper.BindPFlag("verify", joinCmd.Flags().Lookup("verify")); err != nil {
		panic(err)
	}
}
