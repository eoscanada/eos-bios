package bios

import (
	"errors"
	"fmt"
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
	destURL := i.GatewayAddressURL + ref
	req, err := http.NewRequest("GET", destURL, nil)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("Fetching %q from %q...", ref, i.GatewayAddressURL.String())
	resp, err := i.Client.Do(req)
	if err != nil {
		return nil, errors.New("download attempts failed")
	}
	defer resp.Body.Close()

	cnt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		if len(cnt) > 50 {
			cnt = cnt[:50]
		}
		return nil, fmt.Errorf("couldn't get %s, return code: %d, server error: %q", destURL, resp.StatusCode, cnt)
	}
	return cnt, nil
}
