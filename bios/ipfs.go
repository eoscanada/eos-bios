package bios

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	multihash "github.com/multiformats/go-multihash"
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

	if toMultihash(cnt) != ref {
		if len(cnt) > 50 {
			cnt = cnt[:50]
		}
		return nil, fmt.Errorf("contents of %s does not match its hash, content starts with %q, perhaps try a different --ipfs-gateway-address", destURL, string(cnt))
	}

	return cnt, nil
}

func toMultihash(cnt []byte) string {
	hash, _ := multihash.Sum(cnt, multihash.SHA2_256, 32)
	return fmt.Sprintf("/ipfs/%s", hash.B58String())
}
