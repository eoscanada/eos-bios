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
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Version of the software
var Version string

// Flags
var noDiscovery bool
var apiAddress string
var apiAddressURL *url.URL
var ipfsAPIAddress string
var seedNetworkContract = "eosio.disco"

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "eos-bios",
	Short: "A tool to launch EOS.IO Software-based networks",
	Long: `A tool to launch EOS.IO Software-based networks

It provides orchestration of community launches for the mainnet, test
networks, in-house networks as well as local development nodes.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringP("my-discovery", "", "my_discovery_file.yaml", "path to your local discovery file")
	RootCmd.PersistentFlags().StringP("ipfs-gateway-address", "", "https://ipfs.io", "Address to reach an IPFS gateway. Used as a fallback if ipfs-local-gateway-address is unreachable.")
	RootCmd.PersistentFlags().StringP("seednet-api", "", "", "HTTP address of the seed network pointed to by your discovery file")
	RootCmd.PersistentFlags().StringP("seednet-keys", "", "./privkeys.keys", "File containing private keys to your account on the seed network")
	// RootCmd.PersistentFlags().StringVarP(&seedNetworkContract, "seednet-contract", "", "eosio.disco", "Contract account name on the seed network, where to find discovery files from all Block producer candidates.")
	RootCmd.PersistentFlags().StringP("target-api", "", "", "HTTP address to reach the node you are starting (for injection and verification)")

	RootCmd.PersistentFlags().StringP("cache-path", "", ".eos-bios-cache", "directory to store cached data from discovered network")
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose mode.  Causes eos-bios to print debugging messages about its progress.  This is helpful in debugging connection, authentication, and configuration problems.")

	for _, flag := range []string{"cache-path", "my-discovery", "ipfs-gateway-address", "seednet-keys", "seednet-api", "target-api", "verbose"} { // "seednet-contract",
		if err := viper.BindPFlag(flag, RootCmd.PersistentFlags().Lookup(flag)); err != nil {
			panic(err)
		}
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvPrefix("EOS_BIOS")
	viper.AutomaticEnv() // read in environment variables that match
}
