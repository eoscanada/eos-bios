package cmd

import (
	"fmt"

	"github.com/eoscanada/eos-bios/discovery"
	"github.com/spf13/viper"
)

func fetchNetwork(forceFetch bool) (*discovery.Network, error) {
	cachePath := viper.GetString("network.cache_path")
	if cachePath == "" {
		return nil, fmt.Errorf("`network.cache_path` config not specified")
	}

	seedURL := viper.GetString("network.seed_discovery_url")
	if cachePath == "" {
		return nil, fmt.Errorf("`network.seed_discovery_url` config not specified")
	}

	net := discovery.NewNetwork(cachePath, seedURL)

	net.ForceFetch = forceFetch

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
