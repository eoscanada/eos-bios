package bios

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type IPFS struct {
	GatewayAddressURL         *url.URL
	FallbackGatewayAddressURL *url.URL
	Client                    *http.Client
}

func NewIPFS(gatewayAddress, fallbackGatewayAddress string) (out *IPFS, err error) {
	out = &IPFS{
		Client: http.DefaultClient,
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

func (i *IPFS) Get(ref IPFSRef) ([]byte, error) {
	req1, err := http.NewRequest("GET", i.GatewayAddressURL.String()+string(ref), nil)
	if err != nil {
		return nil, err
	}

	req2, err := http.NewRequest("GET", i.FallbackGatewayAddressURL.String()+string(ref), nil)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Fetching %q from primary location (%q)...", ref, i.GatewayAddressURL.String())
	resp, err := i.Client.Do(req1)
	if err != nil {
		fmt.Printf(" failed (%s), trying fallback (%q)...", err, i.FallbackGatewayAddressURL)

		resp, err = i.Client.Do(req2)
		if err != nil {
			fmt.Println(" failed")
			return nil, errors.New("download attempts failed")
		}
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (i *IPFS) GetIPNS(ref IPNSRef) ([]byte, error) {
	return i.Get(IPFSRef(ref))
}
