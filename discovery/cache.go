package discovery

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Network struct {
	cachePath        string
	seedDiscoveryURL string
	visitedURLs      map[string]bool
	discoveredPeers  map[string]*Peer
}

func NewNetwork(cachePath string, seedDiscoveryURL string) *Network {
	return &Network{
		cachePath:        cachePath,
		visitedURLs:      map[string]bool{},
		discoveredPeers:  map[string]*Peer{},
		seedDiscoveryURL: seedDiscoveryURL,
	}
}

func (c *Network) EnsureExists() error {
	return os.MkdirAll(c.cachePath, 0777)
}

func (c *Network) FetchAll() error {
	c.visitedURLs = map[string]bool{}
	c.discoveredPeers = map[string]*Peer{}

	if err := c.FetchOne(c.seedDiscoveryURL); err != nil {
		return fmt.Errorf("fetching %q: %s", c.seedDiscoveryURL, err)
	}

	return nil
}

func (c *Network) FetchOne(discoveryURL string) error {
	if c.visitedURLs[discoveryURL] {
		return nil
	}

	c.visitedURLs[discoveryURL] = true

	disco, err := c.DownloadDiscoveryURL(discoveryURL)
	if err != nil {
		return fmt.Errorf("couldn't download discovery URL: %s", err)
	}

	if (disco.Testnet && disco.Mainnet) || (!disco.Testnet && !disco.Mainnet) {
		return errors.New("mainnet/testnet flag inconsistent, one is require, and only one")
	}

	if disco.EOSIOAccountName == "" {
		return errors.New("eosio_account_name is missing")
	}

	launchData := disco.LaunchData

	// Go through all the things we can download from there
	if err := c.DownloadHashURL(discoveryURL, launchData.BootSequence); err != nil {
		return fmt.Errorf("boot_sequence: %s", err)
	}
	if err := c.DownloadHashURL(discoveryURL, launchData.Snapshot); err != nil {
		return fmt.Errorf("snapshot: %s", err)
	}
	// if err := c.DownloadHashURL(discoveryURL, launchData.SnapshotUnregistered); err != nil {
	// 	return fmt.Errorf("snapshot_unregistered: %s", err)
	// }
	for name, contract := range launchData.Contracts {
		if err := c.DownloadHashURL(discoveryURL, contract.ABI); err != nil {
			return fmt.Errorf("contract %q ABI: %s", name, err)
		}
		if err := c.DownloadHashURL(discoveryURL, contract.Code); err != nil {
			return fmt.Errorf("contract %q Code: %s", name, err)
		}
	}

	c.discoveredPeers[discoveryURL] = &Peer{
		DiscoveryURL: discoveryURL,
		Discovery:    disco,
	}

	// Save the content of the disco file in here

	for _, wingman := range launchData.Wingmen {
		if wingman.Weight > 1.0 || wingman.Weight < 0.0 {
			return fmt.Errorf("weight for wingmen should be between 0.0 and 1.0, %f invalid", wingman.Weight)
		}
		absDiscoURL, err := absoluteURL(discoveryURL, wingman.DiscoveryURL)
		if err != nil {
			return err
		}
		if err := c.FetchOne(absDiscoURL); err != nil {
			return fmt.Errorf("fetching %q: %s", absDiscoURL, err)
		}
	}

	return nil
}

func (c *Network) DownloadHashURL(discoveryURL string, hu HashURL) error {
	if hu.Hash == "" {
		return errors.New("no hash provided")
	}
	if hu.URL == "" {
		return errors.New("no url provided")
	}

	fileName := hu.Hash
	if c.isInCache(fileName) {
		return nil
	}

	destURL, err := absoluteURL(discoveryURL, hu.URL)
	if err != nil {
		return err
	}

	resp, err := http.Get(destURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return err
	}

	hash := sha2(buf.Bytes())

	if hash != hu.Hash {
		return fmt.Errorf("content downloaded from %q does not hash to %q as expected, but hashes to %q", destURL, hu.Hash, hash)
	}

	if err := c.writeToCache(fileName, buf.Bytes()); err != nil {
		return fmt.Errorf("writing %q: %s", fileName, err)
	}

	return nil
}

func (c *Network) writeToCache(fileName string, content []byte) error {
	return ioutil.WriteFile(filepath.Join(c.cachePath, fileName), content, 0666)
}

func (c *Network) isInCache(file string) bool {
	fileName := filepath.Join(c.cachePath, file)

	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
}

func (c *Network) LoadFromCache(initialDiscoveryURL string) error {
	// TODO: start with initialDiscoveryURL
	// read from disk all the BPs, verify the hash data, etc.. ?
	return nil
}

func (c *Network) CalculateWeights() error {
	// build a second map with discoveryURLs alongside account_names...
	for _, peer := range c.discoveredPeers {
		for _, wingman := range peer.Discovery.LaunchData.Wingmen {
			absDiscoURL, err := absoluteURL(peer.DiscoveryURL, wingman.DiscoveryURL)
			if err != nil {
				return err
			}

			wingmanDisco, found := c.discoveredPeers[absDiscoURL]
			if !found {
				return fmt.Errorf("couldn't find %q in list of peers", absDiscoURL)
			}

			addWeight := 1.0
			if wingman.Weight != 0 {
				addWeight = wingman.Weight
			}
			wingmanDisco.TotalWeight += addWeight
		}
	}
	return nil
}

func (c *Network) VerifyGraph() error {
	// Make sure we don't have 2 entities named the same on chain.. EOSIOACcountName being equal.
	seen := map[string]string{}
	for _, peer := range c.discoveredPeers {
		if discoURL := seen[peer.Discovery.EOSIOAccountName]; discoURL != "" {
			return fmt.Errorf("two peers claim the eosio_account_name %q: %q and %q", peer.Discovery.EOSIOAccountName, discoURL, peer.DiscoveryURL)
		}
	}
	return nil
}

func (c *Network) DownloadDiscoveryURL(discoURL string) (out *Discovery, err error) {
	resp, err := http.Get(discoURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return
	}

	fsFile := replaceAllWeirdities(discoURL)
	err = c.writeToCache(fsFile, buf.Bytes())
	if err != nil {
		return
	}

	err = yamlUnmarshal(buf.Bytes(), &out)
	return
}

func sha2(input []byte) string {
	hash := sha256.New()
	_, _ = hash.Write(input) // can't fail
	return hex.EncodeToString(hash.Sum(nil))
}
