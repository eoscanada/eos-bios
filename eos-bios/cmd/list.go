package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the published discovery file for every BP account",
	Run: func(cmd *cobra.Command, args []string) {
		net, err := fetchNetwork(false)
		if err != nil {
			log.Fatalln("fetch network:", err)
		}

		net.PrintDiscoveryFiles(viper.GetBool("verbose"))
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
