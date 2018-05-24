package bios

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/eoscanada/eos-bios/bios/disco"
)

func LoadDiscoveryFromFile(fileName string) (discovery *disco.Discovery, err error) {
	cnt, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	err = yamlUnmarshal(cnt, &discovery)
	if err != nil {
		return
	}

	err = ValidateDiscovery(discovery)
	return
}

func ValidateDiscoveryFile(filename string) error {
	discovery, err := LoadDiscoveryFromFile(filename)
	if err != nil {
		return err
	}

	return ValidateDiscovery(discovery)
}

var p2pAddressRE = regexp.MustCompile(`[a-zA-Z0-9.:-]`)

func ValidateDiscovery(discovery *disco.Discovery) error {
	for _, peer := range discovery.SeedNetworkPeers {
		if peer.Weight > 100 || peer.Weight < 0 {
			return fmt.Errorf("peer %q weight must be between 0 and 100", peer.Account)
		}
	}

	if discovery.TargetAccountName == "" {
		return errors.New("target_account_name is missing")
	}

	if len(discovery.TargetAccountName) != 12 {
		return errors.New("target_account_name should be 12 chars")
	}

	if strings.Contains(string(discovery.TargetAccountName), ".") {
		return errors.New("target_account_name should not contain '.'")
	}

	if strings.Contains(discovery.TargetP2PAddress, "://") {
		return fmt.Errorf("target_p2p_address should be of format ip:port, not prefixed with a protocol")
	}

	if !strings.Contains(discovery.TargetP2PAddress, ":") {
		return errors.New("target_p2p_address doesn't contain port number")
	}

	if !p2pAddressRE.MatchString(discovery.TargetP2PAddress) {
		return errors.New("target_p2p_address is weird, should contain only [a-zA-Z0-9.:-]")
	}

	if !strings.Contains(discovery.TargetHTTPAddress, "://") {
		return fmt.Errorf("target_http_address should include the protocol (like http:// or https://)")
	}

	if strings.Contains(discovery.TargetP2PAddress, " ") {
		return fmt.Errorf("target_p2p_address should not contain spaces")
	}

	if strings.Contains(discovery.TargetHTTPAddress, " ") {
		return fmt.Errorf("target_http_address should not contain spaces")
	}

	// Make sure we have a non-zero weight in the initial authority
	if len(discovery.TargetInitialAuthority.Owner.Keys) == 0 {
		return fmt.Errorf("you need at least one owner key defined in target_initial_authority")
	}

	if len(discovery.TargetInitialAuthority.Active.Keys) == 0 {
		return fmt.Errorf("you need at least one active key defined in target_initial_authority")
	}

	for _, kw := range discovery.TargetInitialAuthority.Owner.Keys {
		if kw.Weight == 0 {
			return fmt.Errorf("weight for owner authority cannot be 0")
		}
	}
	for _, kw := range discovery.TargetInitialAuthority.Active.Keys {
		if kw.Weight == 0 {
			return fmt.Errorf("weight for owner authority cannot be 0")
		}
	}

	// TODO: make sure it's not the DEFAULT key - yeah but that
	// prevents a "clone and boot" scenario.. hmmm.

	// TODO: boot node should orchestrate the PUBLICLY accessible thing..
	//       so we don't need to have two `target_api_address`

	return nil
}
