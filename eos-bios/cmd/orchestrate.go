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
	"net/url"

	bios "github.com/eoscanada/eos-bios"
	eos "github.com/eoscanada/eos-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// orchestrateCmd represents the orchestrate command
var orchestrateCmd = &cobra.Command{
	Use:   "orchestrate",
	Short: "Automate all the operations to launch a new network, by collaborating with other in the launch.",
	Long:  `This operation will auto-select the roles, based on a discovered Network shared amongst participants.`,
	Run: func(cmd *cobra.Command, args []string) {
		// config, err := LoadLocalConfig(*localConfig)
		// if err != nil {
		// 	log.Fatalln("local config load error:", err)
		// }

		net, err := fetchNetwork(true)
		if err != nil {
			log.Fatalln("fetch network:", err)
		}

		apiURL, err := url.Parse(viper.GetString("producer.api_address"))
		if err != nil {
			log.Fatalf("parse config producer.api_address %q: %s\n", viper.GetString("producer.api_address"), err)
		}

		chainID := make([]byte, 32, 32)

		api := eos.New(apiURL, chainID)
		if err != nil {
			log.Fatalln("producer node error:", err)
		}

		api.SetSigner(eos.NewKeyBag())

		config := &bios.Config{}
		config.Producer.MyAccount = viper.GetString("producer.my_account")
		config.Producer.SecretP2PAddress = viper.GetString("producer.secret_p2p_address")

		// Start BIOS
		bios := bios.NewBIOS(net, config, api)

		if err := bios.Init(); err != nil {
			log.Fatalf("BIOS initialization error: %s", err)
		}

		if err := bios.Run(); err != nil {
			log.Fatalf("ERROR RUNNING BIOS: %s", err)
		}

	},
}

func init() {
	RootCmd.AddCommand(orchestrateCmd)
}
