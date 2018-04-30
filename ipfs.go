package bios

import (
	"fmt"
	"net/url"
)

type IPFS struct {
	APIAddressURL             *url.URL
	GatewayAddressURL         *url.URL
	FallbackGatewayAddressURL *url.URL
}

func NewIPFS(apiAddress, gatewayAddress, fallbackGatewayAddress string) (out *IPFS, err error) {
	out = &IPFS{}

	out.APIAddressURL, err = url.Parse(apiAddress)
	if err != nil {
		return nil, fmt.Errorf("parsing api address: %s", err)
	}

	out.GatewayAddressURL, err = url.Parse(gatewayAddress)
	if err != nil {
		return nil, fmt.Errorf("parsing gateway address: %s", err)
	}

	out.FallbackGatewayAddressURL, err = url.Parse(fallbackGatewayAddress)
	if err != nil {
		return nil, fmt.Errorf("parsing fallback gateway address: %s", err)
	}

	return
}
