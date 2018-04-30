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
	publishCmd.PersistentFlags().StringVarP(&ipfsAPIFile, "ipfs-api-file", "", "./.ipfs-data/api", "Pointer to the `api` file in the IPFS_PATH once you started `ipfs daemon`.")

	for _, flag := range []string{"ipfs-api-file"} {
		viper.BindPFlag(flag, publishCmd.Flags().Lookup(flag))
	}
}
