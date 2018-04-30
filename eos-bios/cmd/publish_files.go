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

// filesCmd represents the files command
var filesCmd = &cobra.Command{
	Use:   "files [[file1] file2...]",
	Short: "Publish boot sequence, contracts, snapshot files, etc..",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, ipfs := ipfsClient()

		for _, arg := range args {
			fl, err := os.Open(arg)
			if err != nil {
				fmt.Println("error opening file:", err)
				os.Exit(1)
			}
			defer fl.Close()

			newObj, err := ipfs.Add(fl)
			if err != nil {
				fmt.Println("error adding file:", err)
				os.Exit(1)
			}

			fmt.Println("/ipfs/" + newObj + "   " + arg + " added.")
		}
	},
}

func init() {
	publishCmd.AddCommand(filesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// filesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// filesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
