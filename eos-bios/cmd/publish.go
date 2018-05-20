// Copyright © 2018 Alexandre Bourget <alex@eoscanada.com>

package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/eoscanada/eos-bios/bios/disco"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish my discovery file to the seed network",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		net, err := fetchNetwork(false, false)
		if err != nil {
			log.Fatalln("fetch network:", err)
		}

		launchTime, currentBlock, err := net.LaunchBlockTime(uint32(net.MyPeer.Discovery.SeedNetworkLaunchBlock))
		if err != nil {
			fmt.Println("get last block num failed: ", err)
			os.Exit(1)
		}
		fmt.Printf("Target block: %d (current: %d)\n", net.MyPeer.Discovery.SeedNetworkLaunchBlock, currentBlock)
		past := ""
		if launchTime.Before(time.Now()) {
			past = " - past!"
		}

		fmt.Printf("- Target time: %s (%s%s)\n", humanize.Time(launchTime), launchTime.Format(time.RFC1123Z), past)

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
