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

	"github.com/eoscanada/eos-bios/bios/disco"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// bootCmd represents the boot command
var bootCmd = &cobra.Command{
	Use:   "boot",
	Short: "Triggers hooks to boot a new network or node",
	Long: `This will run the "boot_node" hook with data generated locally for a new network.

The "publish_kickstart_data" will also be run, giving you the opportunity to disseminate what is required for people to join your network.

Boot is what happens when you run "eos-bios orchestrate" and you are selected to be the BIOS Boot node.
`,
	Run: func(cmd *cobra.Command, args []string) {
		net, err := fetchNetwork(viper.GetBool("single"), true)
		if err != nil {
			log.Fatalln("fetch network:", err)
		}

		if viper.GetBool("reset") {
			fmt.Println("Resetting genesis data on seed network")
			_, err := net.SeedNetAPI.SignPushActions(
				disco.NewDeleteGenesis(net.MyPeer.Discovery.SeedNetworkAccountName),
			)
			if err != nil {
				fmt.Println("Error deleting:", err)
				os.Exit(1)
			}
			fmt.Println("Done")
			os.Exit(0)
		}

		b, err := setupBIOS(net)
		if err != nil {
			log.Fatalln("bios setup:", err)
		}

		b.SingleOnly = viper.GetBool("single")
		b.OverrideBootSequenceFile = viper.GetString("override-bootseq")

		if err := b.Init(); err != nil {
			log.Fatalf("BIOS initialization error: %s", err)
		}

		//b.TargetNetAPI.Debug=true

		if err := b.StartBoot(); err != nil {
			log.Fatalf("error booting network: %s", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(bootCmd)

	bootCmd.Flags().BoolP("single", "s", false, "Don't try to discover the world, just boot a local instance.")
	bootCmd.Flags().BoolP("reset", "", false, "Remove the published genesis data from the seed_network, so that others don't accidentally join a defunc or restarted network.")
	bootCmd.Flags().StringP("override-bootseq", "", "", "Override the boot_sequence.yaml file with a local file path (don't used the published one)")

	for _, flag := range []string{"single", "override-bootseq", "reset"} {
		if err := viper.BindPFlag(flag, bootCmd.Flags().Lookup(flag)); err != nil {
			panic(err)
		}
	}
}
