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

	"github.com/eoscanada/eos-bios/bios"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate [some_file.yaml]",
	Short: "Validate check for the integrity of a local discovery file by default, or another file.",
	Long:  "Check your files before you put them out, as to not break the network being crafted.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := viper.GetString("my-discovery")
		if len(args) == 1 {
			filename = args[0]
		}
		if err := bios.ValidateDiscoveryFile(filename); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("File valid:", filename)
	},
}

func init() {
	RootCmd.AddCommand(validateCmd)
}
