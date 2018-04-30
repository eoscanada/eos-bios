package cmd

import (
	"fmt"

	"github.com/eoscanada/eos-bios/discovery"
	"github.com/spf13/viper"
)

func fetchNetwork() (*discovery.Network, error) {
	seedURL := viper.GetString("network.seed_discovery_url")
	if seedURL == "" {
		return nil, fmt.Errorf("`network.seed_discovery_url` config not specified")
	}

	// TODO: use `myDiscoveryFile` instead of a URL
	net := discovery.NewNetwork(cachePath, seedURL)

	net.ForceFetch = !useCache

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
