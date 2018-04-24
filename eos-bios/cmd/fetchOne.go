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

	"github.com/spf13/cobra"
)

// fetchOneCmd represents the fetchOne command
var fetchOneCmd = &cobra.Command{
	Use:   "fetch-one [account_name]",
	Args:  cobra.ExactArgs(1),
	Short: "Refreshes a single account_name, does not attempt to re-download all discovery_urls and traverse the whole graph.",
	Run: func(cmd *cobra.Command, args []string) {
		// Load from cache
		// Then FetchOne with the discovery URL of a given name (that is already in cache?)
		// or take a URL from the command line..
		fmt.Println("fetch-one called")
	},
}

func init() {
	discoveryCmd.AddCommand(fetchOneCmd)
}
