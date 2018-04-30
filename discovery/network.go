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
	"sort"
)

type Network struct {
	ForceFetch bool

	MyPeer *Peer

	cachePath        string
	seedDiscoveryURL string
	visitedURLs      map[string]bool
	discoveredPeers  map[string]*Peer
	orderedPeers     []*Peer
}

func NewNetwork(cachePath string, seedDiscoveryURL string) *Network {
	return &Network{
		cachePath:        cachePath,
		visitedURLs:      map[string]bool{},
		discoveredPeers:  map[string]*Peer{},
		seedDiscoveryURL: seedDiscoveryURL,
	}
}

func (c *Network) ensureExists() error {
	return os.MkdirAll(c.cachePath, 0777)
}

func (c *Network) FetchAll() error {
	c.visitedURLs = map[string]bool{}
	c.discoveredPeers = map[string]*Peer{}

	err := c.ensureExists()
	if err != nil {
		return fmt.Errorf("error creating cache path: %s", err)
	}

	if err := c.FetchOne(c.seedDiscoveryURL); err != nil {
		return fmt.Errorf("fetching %q: %s", c.seedDiscoveryURL, err)
	}

	return nil
}

func (c *Network) ValidateLocalFile(filename string) error {
	// simulate DownloadDiscoveryURL with a local file, and run all
	// the validation we have from `if disco.Testnet && disco.Mainnet`,
	// etc..

	return nil
}

func (c *Network) ChainID() []byte {
	// TODO: compute based on all the hashes in the elected launchdata?
	// have a value be voted for ?
	return make([]byte, 32, 32)
}

func (c *Network) FetchOne(discoveryURL string) error {
	if c.visitedURLs[discoveryURL] {
		return nil
	}

	c.visitedURLs[discoveryURL] = true

	disco, rawDisco, err := c.DownloadDiscoveryURL(discoveryURL)
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
	fsFile := replaceAllWeirdities(discoveryURL)
	// fmt.Printf("Discovery: writing %q to cache\n", fsFile)
	err = c.writeToCache(fsFile, rawDisco)
	if err != nil {
		return fmt.Errorf("writing discovery data to %q: %s", fsFile, err)
	}

	fmt.Printf("Discovery: traversing %d peers\n", len(launchData.Peers))
	for _, peerLink := range launchData.Peers {
		if peerLink.Weight > 1.0 || peerLink.Weight < 0.0 {
			return fmt.Errorf("weight for peers should be between 0.0 and 1.0, %f invalid", peerLink.Weight)
		}
		absDiscoURL, err := absoluteURL(discoveryURL, peerLink.DiscoveryURL)
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
		// fmt.Printf("Discovery: %q in cache\n", hu.Hash)
		return nil
	}

	destURL, err := absoluteURL(discoveryURL, hu.URL)
	if err != nil {
		return err
	}

	fmt.Printf("Discovery: downloading %q from %q\n", hu.Hash, destURL)
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

func (c *Network) ReadFromCache(fileName string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(c.cachePath, fileName))
}

func (c *Network) ReaderFromCache(fileName string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(c.cachePath, fileName))
}

func (c *Network) FileNameFromCache(fileName string) string {
	return filepath.Join(c.cachePath, fileName)
}

func (c *Network) LoadCache(initialDiscoveryURL string) error {
	// TODO: start with initialDiscoveryURL
	// read from disk all the BPs, verify the hash data, etc.. ?
	return nil
}

func (c *Network) CalculateWeights() error {
	// build a second map with discoveryURLs alongside account_names...
	var allPeers []*Peer
	for _, peer := range c.discoveredPeers {
		for _, peerLink := range peer.Discovery.LaunchData.Peers {
			absDiscoURL, err := absoluteURL(peer.DiscoveryURL, peerLink.DiscoveryURL)
			if err != nil {
				return err
			}

			if peer.DiscoveryURL == absDiscoURL {
				// Can't vouch for yourself
				continue
			}

			peerLinkDisco, found := c.discoveredPeers[absDiscoURL]
			if !found {
				return fmt.Errorf("couldn't find %q in list of peers", absDiscoURL)
			}

			addWeight := 1.0
			if peerLink.Weight != 0 {
				addWeight = peerLink.Weight
			}
			peerLinkDisco.TotalWeight += addWeight
		}

		allPeers = append(allPeers, peer)
	}

	// Sort the `orderedPeers`
	sort.Slice(allPeers, func(i, j int) bool {
		if allPeers[i].TotalWeight == allPeers[j].TotalWeight {
			return allPeers[i].DiscoveryURL < allPeers[j].DiscoveryURL
		}
		return allPeers[i].TotalWeight > allPeers[j].TotalWeight
	})

	c.orderedPeers = allPeers

	return nil
}

func (c *Network) OrderedPeers() []*Peer {
	return c.orderedPeers
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

func (c *Network) DownloadDiscoveryURL(discoveryURL string) (out *Discovery, rawDiscovery []byte, err error) {
	var buf bytes.Buffer

	fsFile := replaceAllWeirdities(discoveryURL)
	if c.ForceFetch || !c.isInCache(fsFile) {
		fmt.Println("Discovery: downloading discovery_url", discoveryURL)
		resp, err := http.Get(discoveryURL)
		if err != nil {
			return out, nil, err
		}
		defer resp.Body.Close()

		_, err = io.Copy(&buf, resp.Body)
		if err != nil {
			return out, nil, err
		}
	} else {
		cnt, err := c.ReadFromCache(fsFile)
		if err != nil {
			return out, cnt, err
		}

		_, _ = buf.Write(cnt)
	}

	rawDiscovery = buf.Bytes()

	err = yamlUnmarshal(rawDiscovery, &out)
	return
}

func sha2(input []byte) string {
	hash := sha256.New()
	_, _ = hash.Write(input) // can't fail
	return hex.EncodeToString(hash.Sum(nil))
}

func (c *Network) PrintOrderedPeers() {
	fmt.Println("###############################################################################################")
	fmt.Println("####################################    PEER NETWORK    #######################################")
	fmt.Println("")

	fmt.Printf("BIOS NODE: %s\n", c.orderedPeers[0].String())
	for i := 1; i < 22 && len(c.orderedPeers) > i; i++ {
		fmt.Printf("ABP %02d:    %s\n", i, c.orderedPeers[i].String())
	}
	for i := 22; len(c.orderedPeers) > i; i++ {
		fmt.Printf("Part. %02d:  %s\n", i, c.orderedPeers[i].String())
	}
	fmt.Println("")
	fmt.Println("###############################################################################################")
	fmt.Println("")
}

// ReachedConsensus reads all the hashes of the top-level peers and
// returns true if we have reached an agreement on the content to
// inject in the chain.
func (c *Network) ReachedConsensus() bool {
	// TODO: Implement the logic that determines the consensus.. right
	// now it's just the weights in order.. and the top-most wins: we use
	// its configuration.
	return true
}

func (c *Network) ConsensusLaunchData() (*LaunchData, error) {
	// TODO: implement the algo to create a Discovery file based on
	// the most vouched for hashes for all the components.
	//
	// Will that work ? Will that make sense ?
	//
	// Cycle through the top peers, take the most vetted
	return &(c.orderedPeers[0].Discovery.LaunchData), nil
}
