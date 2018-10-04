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

// bootCmd represents the boot command
var bootCmd = &cobra.Command{
	Use:   "boot [boot_sequence.yaml]",
	Short: "Boots a new nodeos and injects the boot sequence.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		b, err := setupBIOS()
		if err != nil {
			log.Fatalln("bios setup:", err)
		}

		if len(args) == 0 {
			b.BootSequenceFile = "boot_sequence.yaml"
		} else {
			b.BootSequenceFile = args[0]
		}

		b.ReuseGenesis = viper.GetBool("reuse-genesis")

		if err := b.Boot(); err != nil {
			log.Fatalf("BIOS boot error: %s", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(bootCmd)

	bootCmd.Flags().BoolP("reuse-genesis", "", false, "Re-load genesis data from genesis.json, genesis.pub and genesis.key instead of creating a new one.")

	for _, flag := range []string{"reuse-genesis"} {
		if err := viper.BindPFlag(flag, bootCmd.Flags().Lookup(flag)); err != nil {
			panic(err)
		}
	}
}
