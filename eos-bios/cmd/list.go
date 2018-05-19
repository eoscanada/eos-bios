package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List shows all the networks discovered from the source contract, not only the one you are participating with.",
	Run: func(cmd *cobra.Command, args []string) {
		net, err := fetchNetwork(false, false)
		if err != nil {
			log.Fatalln("fetch network:", err)
		}

		net.ListNetworks(viper.GetBool("verbose"))
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
