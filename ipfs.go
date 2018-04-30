package bios

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type IPFS struct {
	APIAddressURL             *url.URL
	GatewayAddressURL         *url.URL
	FallbackGatewayAddressURL *url.URL
	Client                    *http.Client
}

func NewIPFS(apiAddress, gatewayAddress, fallbackGatewayAddress string) (out *IPFS, err error) {
	out = &IPFS{
		Client: http.DefaultClient,
	}

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

func (i *IPFS) Get(ref IPFSRef) ([]byte, error) {
	req1, err := http.NewRequest("GET", i.GatewayAddressURL.String()+string(ref), nil)
	if err != nil {
		return nil, err
	}

	req2, err := http.NewRequest("GET", i.FallbackGatewayAddressURL.String()+string(ref), nil)
	if err != nil {
		return nil, err
	}

	reqs := []*http.Request{req1, req2}
	var resp *http.Response
	for _, req := range reqs {
		resp, err = i.Client.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "NOTE: %q unavailable (%s), trying fallback\n", req.URL.String(), err)
			continue
		}
	}
	if err != nil {
		return nil, fmt.Errorf("gateway reqs failed: %s", err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (i *IPFS) GetIPNS(ref IPNSRef) ([]byte, error) {
	return i.Get(IPFSRef(ref))
}
