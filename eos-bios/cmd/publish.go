// Copyright © 2018 Alexandre Bourget <alex@eoscanada.com>

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish some content to IPFS for others to discover.",
	Long:  ``,
}

func init() {
	RootCmd.AddCommand(publishCmd)
	publishCmd.PersistentFlags().StringVarP(&ipfsAPIAddress, "ipfs-api-address", "", "http://127.0.0.1:5001", "address to reach the API of a local ipfs node")

	for _, flag := range []string{"ipfs-api-address"} {
		viper.BindPFlag(flag, publishCmd.Flags().Lookup(flag))
	}
}
