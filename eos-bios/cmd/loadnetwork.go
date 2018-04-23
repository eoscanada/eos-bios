package cmd

import (
	"fmt"

	"github.com/eoscanada/eos-bios/discovery"
	"github.com/spf13/viper"
)

func fetchNetwork(fromCache bool) (*discovery.Network, error) {
	cachePath := viper.GetString("network.cache_path")
	if cachePath == "" {
		return nil, fmt.Errorf("Please specify a `cache_path` under `network` in your configuration.")
	}

	seedURL := viper.GetString("network.seed_discovery_url")
	if cachePath == "" {
		return nil, fmt.Errorf("Please specify a `seed_discovery_url` under `network` in your configuration.")
	}

	net := discovery.NewNetwork(cachePath, seedURL)

	if err := net.EnsureExists(); err != nil {
		return nil, fmt.Errorf("Error creating cache path: %s", err)
	}

	if err := net.FetchAll(); err != nil {
		return nil, fmt.Errorf("fetch-all error: %s", err)
	}

	if err := net.VerifyGraph(); err != nil {
		return nil, fmt.Errorf("graph inconsistent: %s", err)
	}

	if err := net.CalculateWeights(); err != nil {
		return nil, fmt.Errorf("error calculating weights: %s", err)
	}

	return net, nil
}
