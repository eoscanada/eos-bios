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

var discoveryCmd = &cobra.Command{
	Use:   "discovery",
	Short: "Publish the discovery file specified by `--my-discovery` to IPFS.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		info, ipfs := ipfsClient()

		fmt.Printf("Reading discovery file... ")
		fl, err := os.Open(myDiscoveryFile)
		if err != nil {
			fmt.Println("failed")
			fmt.Fprintf(os.Stderr, "error opening %q: %s\n", myDiscoveryFile, err)
			os.Exit(1)
		}
		defer fl.Close()
		fmt.Println("ok")

		fmt.Printf("Publishing discovery file... ")
		newObj, err := ipfs.Add(fl)
		if err != nil {
			fmt.Println("failed")
			fmt.Fprintln(os.Stderr, "error adding content to ipfs:", err)
			os.Exit(1)
		}
		fmt.Println("/ipfs/" + newObj + " published")

		fmt.Printf("Updating IPNS link /ipns/" + info.ID + " ... ")
		if err = ipfs.Publish("", newObj); err != nil {
			fmt.Println("failed")
			fmt.Fprintln(os.Stderr, "error publishing new ipns address:", err)
			os.Exit(1)
		}
		fmt.Println("ok")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("You can now provide your IPNS link to your network in this form:")
		fmt.Println("")
		fmt.Println("    /ipns/" + info.ID)
		fmt.Println("")
		// fmt.Println("Alternatively, you can add/update a TXT record to your domain with this data:")
		// fmt.Println("")
		// fmt.Println(`    "dnslink=/ipfs/` + newObj + `"`)
		// fmt.Println("")
		// fmt.Println("and then provide to your network a link such as:")
		// fmt.Println("")
		// fmt.Println("    /ipns/testnet-myname.eosantarcti.ca")
		// fmt.Println("")
		// fmt.Println("which will resolve to this document.")
	},
}

func init() {
	publishCmd.AddCommand(discoveryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// discoveryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// discoveryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
