package bios

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type IPFS struct {
	GatewayAddressURL string
	Client            *http.Client
}

func NewIPFS(gatewayAddress string) (out *IPFS) {
	out = &IPFS{
		Client:            http.DefaultClient,
		GatewayAddressURL: gatewayAddress,
	}
	return
}

func (i *IPFS) Get(ref string) ([]byte, error) {
	req, err := http.NewRequest("GET", i.GatewayAddressURL+ref, nil)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("Fetching %q from %q...", ref, i.GatewayAddressURL.String())
	resp, err := i.Client.Do(req)
	if err != nil {
		return nil, errors.New("download attempts failed")
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
