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
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
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
	Short: "A tool to launch EOS.IO Software-based networks and envs",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	homedir, err := homedir.Dir()
	if err != nil {
		fmt.Println("couldn't find home dir:", err)
		os.Exit(1)
	}

	RootCmd.PersistentFlags().StringP("api-url", "", "http://localhost:8888", "HTTP address to reach the node you are starting (for injection and validation)")
	RootCmd.PersistentFlags().BoolP("hack-voting-accounts", "", false, "This will take accounts with large stakes and put a well known public key in place, so the community can test voting.")

	RootCmd.PersistentFlags().BoolP("write-actions", "", false, "Write actions to actions.jsonl upon join or boot")
	RootCmd.PersistentFlags().StringP("cache-path", "", filepath.Join(homedir, ".eos-bios-cache"), "directory to store cached data from discovered network")
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "Display verbose output (also see 'output.log')")

	for _, flag := range []string{"cache-path", "write-actions", "api-url", "verbose", "hack-voting-accounts"} {
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
