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

	"github.com/spf13/cobra"
)

// fetchAllCmd represents the fetchAll command
var fetchAllCmd = &cobra.Command{
	Use:   "fetch-all",
	Short: "Discover and update info about all peers in the network, based on an initial discovery URL.",
	Long:  `This uses the "network.seed_discovery_url" key in your configuration to start discovery.`,
	Run: func(cmd *cobra.Command, args []string) {
		// load cache_path
		// the second param to `NewNetwork` can be the path where this file will live,
		// for references
		net, err := fetchNetwork(false)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Fetched successfully")

		// TODO: display the orderedPeers
		net.PrintOrderedPeers()
	},
}

func init() {
	discoveryCmd.AddCommand(fetchAllCmd)
}
