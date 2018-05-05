package bios

import (
	"encoding/json"
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

	for _, peer := range disco.LaunchData.Peers {
		if !strings.HasPrefix(string(peer.DiscoveryLink), "/ipns/") {
			return fmt.Errorf("peer link %q invalid, should start with '/ipns/'", peer.DiscoveryLink)
		}
		if peer.Weight > 1.0 {
			return fmt.Errorf("peer %q weight must be between 0.0 and 1.0", peer.DiscoveryLink)
		}
	}

	return nil
}
