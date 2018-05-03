package bios

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	multihash "github.com/multiformats/go-multihash"
)

type Network struct {
	ForceFetch bool

	MyPeer *Peer

	IPFS *IPFS

	cachePath       string
	myDiscoveryFile string
	discoveredIPNS  map[IPNSRef]bool
	discoveredPeers map[IPFSRef]*Peer
	orderedPeers    []*Peer

	lastFetch time.Time
}

func NewNetwork(cachePath string, myDiscoveryFile string, ipfs *IPFS) *Network {
	return &Network{
		IPFS:            ipfs,
		cachePath:       cachePath,
		myDiscoveryFile: myDiscoveryFile,
	}
}

func (c *Network) ensureExists() error {
	return os.MkdirAll(c.cachePath, 0777)
}

func (net *Network) UpdateGraph() error {
	if time.Now().Before(net.lastFetch.Add(2 * time.Minute)) {
		return nil
	}

	if err := net.traverseGraph(); err != nil {
		return fmt.Errorf("traversing graph: %s", err)
	}

	if err := net.verifyGraph(); err != nil {
		return fmt.Errorf("verifying graph: %s", err)
	}

	if err := net.calculateWeights(); err != nil {
		return fmt.Errorf("calculating weights: %s", err)
	}

	return nil
}

func (c *Network) traverseGraph() error {
	c.discoveredIPNS = map[IPNSRef]bool{}
	c.discoveredPeers = map[IPFSRef]*Peer{}

	err := c.ensureExists()
	if err != nil {
		return fmt.Errorf("error creating cache path: %s", err)
	}

	//fmt.Println("Cache ready")

	// TODO: how do we handle when someone points to *us* ? We should
	// have a way to find our canonical URL..
	rawDisco, err := ioutil.ReadFile(c.myDiscoveryFile)
	if err != nil {
		return err
	}

	var disco *Discovery
	err = yamlUnmarshal(rawDisco, &disco)
	if err != nil {
		return err
	}

	ipfsRef := toMultihash(rawDisco)

	c.MyPeer = &Peer{
		Discovery: disco,
	}

	return c.traversePeer(disco, IPNSRef("local "+c.myDiscoveryFile), ipfsRef)
}

func (c *Network) fetchOne(peerLink *PeerLink) error {
	if c.discoveredIPNS[peerLink.DiscoveryLink] {
		fmt.Printf("    - traversed already!\n")
		return nil
	}

	disco, rawDisco, err := c.FetchDiscoveryLink(peerLink)
	if err != nil {
		return fmt.Errorf("couldn't download discovery URL: %s", err)
	}

	if disco.EOSIOAccountName == c.MyPeer.Discovery.EOSIOAccountName {
		fmt.Printf("    - was myself!\n")
		return nil
	}
	fmt.Printf("    - %q (%q)\n", disco.EOSIOAccountName, disco.OrganizationName)

	ipfsRef := toMultihash(rawDisco)

	c.discoveredIPNS[peerLink.DiscoveryLink] = true

	peerLink.resolvedRef = ipfsRef

	return c.traversePeer(disco, peerLink.DiscoveryLink, ipfsRef)
}

func toMultihash(cnt []byte) IPFSRef {
	hash, _ := multihash.Sum(cnt, multihash.SHA2_256, 32)
	return IPFSRef(fmt.Sprintf("/ipfs/%s", hash.B58String()))
}

func (c *Network) traversePeer(disco *Discovery, ipnsRef IPNSRef, ipfsRef IPFSRef) error {
	fmt.Printf("Loading launch data from %q (%q, %s)...\n", disco.EOSIOAccountName, disco.OrganizationName, ipnsRef)
	if (disco.Testnet && disco.Mainnet) || (!disco.Testnet && !disco.Mainnet) {
		return errors.New("mainnet/testnet flag inconsistent, one is require, and only one")
	}

	if disco.EOSIOAccountName == "" {
		return errors.New("eosio_account_name is missing")
	}

	launchData := disco.LaunchData

	// Go through all the things we can download from there
	if err := c.DownloadIPFSRef(launchData.BootSequence); err != nil {
		return fmt.Errorf("boot_sequence: %s", err)
	}
	if err := c.DownloadIPFSRef(launchData.Snapshot); err != nil {
		return fmt.Errorf("snapshot: %s", err)
	}
	// if err := c.DownloadHashURL(discoveryURL, launchData.SnapshotUnregistered); err != nil {
	// 	return fmt.Errorf("snapshot_unregistered: %s", err)
	// }
	for name, contract := range launchData.Contracts {
		if err := c.DownloadIPFSRef(contract.ABI); err != nil {
			return fmt.Errorf("contract %q ABI: %s", name, err)
		}
		if err := c.DownloadIPFSRef(contract.Code); err != nil {
			return fmt.Errorf("contract %q Code: %s", name, err)
		}
	}

	c.discoveredPeers[ipfsRef] = &Peer{
		DiscoveryFile: ipfsRef,
		DiscoveryLink: ipnsRef,
		Discovery:     disco,
	}

	fmt.Printf("- has %d peer(s)\n", len(launchData.Peers))

	for _, peerLink := range launchData.Peers {
		if peerLink.Weight > 1.0 || peerLink.Weight < 0.0 {
			fmt.Printf("WARN: peer %q weight not between 0.0 and 1.0, not including in graph\n", peerLink.DiscoveryLink)
			continue
		}

		if err := c.fetchOne(peerLink); err != nil {
			return fmt.Errorf("fetching %q: %s", peerLink.DiscoveryLink, err)
		}
	}

	return nil
}

func (c *Network) FetchDiscoveryLink(peerLink *PeerLink) (out *Discovery, rawDiscovery []byte, err error) {
	// Resolve recursive the discoveryLink (through /ipns, then /ipfs/Qm.../path to /ipfs/Qmcontent)

	fmt.Printf("  - peer %s comment=%q, weight=%.2f\n", peerLink.DiscoveryLink, peerLink.Comment, peerLink.Weight)
	rawDiscovery, err = c.IPFS.GetIPNS(peerLink.DiscoveryLink)
	if err != nil {
		return
	}

	err = yamlUnmarshal(rawDiscovery, &out)
	return
}

func (c *Network) DownloadIPFSRef(ref IPFSRef) error {
	if ref == "" {
		return errors.New("no hash provided")
	}
	if !strings.HasPrefix(string(ref), "/ipfs/") {
		return fmt.Errorf("ipfs ref should start with'/ipfs/': %q", ref)
	}

	if c.isInCache(ref) {
		//fmt.Printf("ipfs ref: %q in cache\n", ref)
		return nil
	}

	cnt, err := c.IPFS.Get(ref)
	if err != nil {
		return err
	}

	if err := c.writeToCache(ref, cnt); err != nil {
		return err
	}

	return nil
}

func (c *Network) writeToCache(ref IPFSRef, content []byte) error {
	fileName := replaceAllWeirdities(string(ref))
	return ioutil.WriteFile(filepath.Join(c.cachePath, fileName), content, 0666)
}

func (c *Network) isInCache(ref IPFSRef) bool {
	fileName := filepath.Join(c.cachePath, replaceAllWeirdities(string(ref)))

	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
}

func (c *Network) ReadFromCache(ref IPFSRef) ([]byte, error) {
	fileName := replaceAllWeirdities(string(ref))
	return ioutil.ReadFile(filepath.Join(c.cachePath, fileName))
}

func (c *Network) ReaderFromCache(ref IPFSRef) (io.ReadCloser, error) {
	fileName := replaceAllWeirdities(string(ref))
	return os.Open(filepath.Join(c.cachePath, fileName))
}

func (c *Network) FileNameFromCache(ref IPFSRef) string {
	fileName := replaceAllWeirdities(string(ref))
	return filepath.Join(c.cachePath, fileName)
}

func (c *Network) LoadCache(initialDiscoveryURL string) error {
	// TODO: start with initialDiscoveryURL
	// read from disk all the BPs, verify the hash data, etc.. ?
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

func (c *Network) calculateWeights() error {
	// build a second map with discoveryURLs alongside account_names...
	var allPeers []*Peer
	for _, peer := range c.discoveredPeers {
		for _, peerLink := range peer.Discovery.LaunchData.Peers {
			// TODO: double-check this.. wuuut
			if peer.DiscoveryFile == peerLink.resolvedRef {
				// Can't vouch for yourself
				continue
			}

			peerLinkDisco, found := c.discoveredPeers[peerLink.resolvedRef]
			if !found {
				return fmt.Errorf("couldn't find %q (resolved peer ref) in list of peers", peerLink.resolvedRef)
			}

			if peer.Discovery.EOSIOAccountName == peerLinkDisco.Discovery.EOSIOAccountName {
				// hmm.. don't count weight on your own account..
				continue
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
			return allPeers[i].DiscoveryFile < allPeers[j].DiscoveryFile
		}
		return allPeers[i].TotalWeight > allPeers[j].TotalWeight
	})

	c.orderedPeers = allPeers

	return nil
}

func (c *Network) OrderedPeers() []*Peer {
	return c.orderedPeers
}

func (c *Network) verifyGraph() error {
	// Make sure we don't have 2 entities named the same on chain.. EOSIOACcountName being equal.
	seen := map[string]string{}
	for _, peer := range c.discoveredPeers {
		if discoURL := seen[peer.Discovery.EOSIOAccountName]; discoURL != "" {
			return fmt.Errorf("two peers claim the eosio_account_name %q: %q and %q", peer.Discovery.EOSIOAccountName, discoURL, peer.DiscoveryFile)
		}
	}
	return nil
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
