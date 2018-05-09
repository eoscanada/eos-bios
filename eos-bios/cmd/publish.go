// Copyright © 2018 Alexandre Bourget <alex@eoscanada.com>

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/eoscanada/eos-bios/bios/disco"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish my discovery file to the seed network",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		net, err := fetchNetwork(false)
		if err != nil {
			log.Fatalln("fetch network:", err)
		}

		_, err = net.SeedNetAPI.SignPushActions(
			disco.NewUpdateDiscovery(net.MyPeer.Discovery.SeedNetworkAccountName, net.MyPeer.Discovery),
		)

		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}

		fmt.Println("Done.")
	},
}

func init() {
	RootCmd.AddCommand(publishCmd)
}
