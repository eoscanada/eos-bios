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

// orchestrateCmd represents the orchestrate command
var orchestrateCmd = &cobra.Command{
	Use:   "orchestrate",
	Short: "Automate all the operations to launch a new network, by collaborating with other in the launch.",
	Long:  `This operation will auto-select the roles, based on a discovered Network shared amongst participants.`,
	Run: func(cmd *cobra.Command, args []string) {
		net, err := fetchNetwork(false, true)
		if err != nil {
			log.Fatalln("fetch network:", err)
		}

		if elect := viper.GetString("elect"); elect != "" {
			net.CalculateNetworkWeights(elect)
		}

		b, err := setupBIOS(net)
		if err != nil {
			log.Fatalln("bios setup:", err)
		}

		if err := b.Init(); err != nil {
			log.Fatalf("BIOS initialization error: %s", err)
		}

		if err := b.StartOrchestrate(); err != nil {
			log.Fatalf("error orchestrating: %s", err)
		}

	},
}

func init() {
	RootCmd.AddCommand(orchestrateCmd)

	// orchestrateCmd.Flags().String("elect", "", "Force the election of the given BIOS Boot node")

	// if err := viper.BindPFlag("elect", orchestrateCmd.Flags().Lookup("elect")); err != nil {
	// 	panic(err)
	// }
}
