package cmd

import (
	"fmt"
)

func fetchNetwork() (*Network, error) {
	net := NewNetwork(cachePath, myDiscoveryFile)

	net.ForceFetch = !noDiscovery

	if err := net.TraverseGraph(); err != nil {
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
