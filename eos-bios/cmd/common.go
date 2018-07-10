package cmd

import (
	"github.com/eoscanada/eos-bios/bios"
	"github.com/eoscanada/eos-go"
	"github.com/spf13/viper"
)

func setupBIOS() (b *bios.BIOS, err error) {
	targetNetHTTP := viper.GetString("api-url")
	targetNetAPI := eos.New(targetNetHTTP)
	targetNetAPI.SetSigner(eos.NewKeyBag())

	logger := bios.NewLogger()
	logger.Debug = viper.GetBool("verbose")

	b = bios.NewBIOS(logger, viper.GetString("cache-path"), targetNetAPI)
	b.WriteActions = viper.GetBool("write-actions")
	b.HackVotingAccounts = viper.GetBool("hack-voting-accounts")
	return b, nil
}
