package bios

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	yaml2json "github.com/bronze1man/go-yaml2json"
)

func yamlUnmarshal(cnt []byte, v interface{}) error {
	jsonCnt, err := yaml2json.Convert(cnt)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonCnt, v)
}

func ValidateDiscoveryFile(filename string) error {
	cnt, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var disco *Discovery
	err = yamlUnmarshal(cnt, &disco)
	if err != nil {
		return err
	}

	return ValidateDiscovery(disco)
}

func ValidateDiscovery(disco *Discovery) error {
	for _, peer := range disco.LaunchData.Peers {
		if !strings.HasPrefix(string(peer.DiscoveryLink), "/ipns/") {
			return fmt.Errorf("peer link %q invalid, should start with '/ipns/'", peer.DiscoveryLink)
		}
		if peer.Weight > 1.0 {
			return fmt.Errorf("peer %q weight must be between 0.0 and 1.0", peer.DiscoveryLink)
		}
	}

	if (disco.Testnet && disco.Mainnet) || (!disco.Testnet && !disco.Mainnet) {
		return errors.New("mainnet/testnet flag inconsistent, one is require, and only one")
	}

	if disco.EOSIOAccountName == "" {
		return errors.New("eosio_account_name is missing")
	}

	// IMPORTANT for USEABILITY:

	// TODO: Validate the `eosio_p2p` is the right format, with a port.
	// Prevent `127.0.0.1`, `localhost`, and `192.168` and local stuff ?

	// TODO: check that p2p nodes don't end with `example.com` or something..

	// TODO: check that `eosio_http` has `http://` prefix, not necessarily a port if standard.
	// TODO: check that `eosio_https` has `https://` prefix.

	// launch ethereum block present.. within reasonable boundaries
	//

	return nil
}
