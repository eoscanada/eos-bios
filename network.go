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

	"github.com/eoscanada/eos-bios/disco"
	"github.com/eoscanada/eos-go"
	"github.com/ryanuber/columnize"
)

type Network struct {
	MyPeer *Peer

	seedNetAPI      *eos.API
	seedNetContract string

	ipfs            *IPFS
	cachePath       string
	discoveredPeers map[eos.AccountName]*Peer
	orderedPeers    []*Peer
	candidates      map[string]*disco.Discovery
}

func NewNetwork(cachePath string, myDiscovery *disco.Discovery, ipfs *IPFS, seedNetContract string, seedNetAPI *eos.API) *Network {
	return &Network{
		seedNetAPI:      seedNetAPI,
		ipfs:            ipfs,
		cachePath:       cachePath,
		seedNetContract: seedNetContract,
		MyPeer: &Peer{
			Discovery: myDiscovery,
		},
	}
}

func (net *Network) ensureExists() error {
	return os.MkdirAll(net.cachePath, 0777)
}

func (net *Network) UpdateGraph() error {
	net.discoveredPeers = map[eos.AccountName]*Peer{}
	net.candidates = make(map[string]*disco.Discovery)

	rows, err := net.seedNetAPI.GetTableRows(
		eos.GetTableRowsRequest{
			JSON:     true,
			Scope:    "eosio.disco",
			Code:     "eosio.disco",
			Table:    "discovery",
			TableKey: "id",
			//LowerBound: "",
			//UpperBound: "",
			Limit: 1000,
		},
	)
	if err != nil {
		return fmt.Errorf("get table rows: %s", err)
	}

	var candidates []*disco.Discovery
	if err := rows.JSONToStructs(&candidates); err != nil {
		return fmt.Errorf("reading discovery from table: %s", err)
	}

	for _, cand := range candidates {
		// verify cand discovery
		net.candidates[string(cand.SeedNetworkAccountName)] = cand
	}

	if err = net.ensureExists(); err != nil {
		return fmt.Errorf("error creating cache path: %s", err)
	}

	if err := net.traversePeer(net.MyPeer.Discovery); err != nil {
		return fmt.Errorf("traversing graph: %s", err)
	}

	if err := net.calculateWeights(); err != nil {
		return fmt.Errorf("calculating weights: %s", err)
	}

	return nil
}

func (net *Network) traversePeer(discoFile *disco.Discovery) error {
	// TODO: should we simply remove the peer if its discovery file is invalid ?
	// flag it as such? and zero its weights everywhere ?
	if err := ValidateDiscovery(discoFile); err != nil {
		return err
	}

	net.discoveredPeers[discoFile.SeedNetworkAccountName] = &Peer{Discovery: discoFile}

	for _, contentRef := range discoFile.TargetContents {
		if err := net.DownloadIPFSRef(contentRef.Ref); err != nil {
			return fmt.Errorf("content %q: %s", contentRef.Name, err)
		}
	}

	fmt.Printf("- has %d peer(s)\n", len(discoFile.SeedNetworkPeers))

	for _, peerLink := range discoFile.SeedNetworkPeers {
		fmt.Printf("  - peer %s comment=%q, weight=%.2f\n", peerLink.Account, peerLink.Comment, peerLink.Weight)

		peerDisco, found := net.candidates[string(peerLink.Account)]
		if !found {
			fmt.Println("    - peer not found")
			continue
		}

		if net.discoveredPeers[peerDisco.SeedNetworkAccountName] != nil {
			fmt.Printf("    - already added %q\n", peerDisco.SeedNetworkAccountName)
			return nil
		}

		fmt.Printf("    - adding %q\n", peerDisco.SeedNetworkAccountName)

		if err := net.traversePeer(peerDisco); err != nil {
			return err
		}
	}

	return nil
}

func (net *Network) DownloadIPFSRef(ref string) error {
	if ref == "" {
		return errors.New("no hash provided")
	}
	if !strings.HasPrefix(ref, "/ipfs/") {
		return fmt.Errorf("ipfs ref should start with'/ipfs/': %q", ref)
	}

	if net.isInCache(string(ref)) {
		//fmt.Printf("ipfs ref: %q in cache\n", ref)
		return nil
	}

	cnt, err := net.ipfs.Get(IPFSRef(ref))
	if err != nil {
		return err
	}

	if err := net.writeToCache(string(ref), cnt); err != nil {
		return err
	}

	return nil
}

func (net *Network) writeToCache(ref string, content []byte) error {
	fileName := replaceAllWeirdities(ref)
	return ioutil.WriteFile(filepath.Join(net.cachePath, fileName), content, 0666)
}

func (net *Network) isInCache(ref string) bool {
	fileName := filepath.Join(net.cachePath, replaceAllWeirdities(string(ref)))

	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
}

func (net *Network) ReadFromCache(ref string) ([]byte, error) {
	fileName := replaceAllWeirdities(ref)
	return ioutil.ReadFile(filepath.Join(net.cachePath, fileName))
}

func (net *Network) ReaderFromCache(ref IPFSRef) (io.ReadCloser, error) {
	fileName := replaceAllWeirdities(string(ref))
	return os.Open(filepath.Join(net.cachePath, fileName))
}

func (net *Network) FileNameFromCache(ref IPFSRef) string {
	fileName := replaceAllWeirdities(string(ref))
	return filepath.Join(net.cachePath, fileName)
}

func (net *Network) ChainID() []byte {
	// TODO: compute based on all the hashes in the elected launchdata?
	// have a value be voted for ?
	return make([]byte, 32, 32)
}

func (net *Network) calculateWeights() error {
	// build a second map with discoveryURLs alongside account_names...
	var allPeers []*Peer
	for _, peer := range net.discoveredPeers {

		fmt.Println("First level peer", peer.DiscoveryLink)
		for _, peerLink := range peer.Discovery.SeedNetworkPeers {

			if peer.Discovery.SeedNetworkAccountName == peerLink.Account {
				// Can't vote for yourself
				continue
			}

			peerLinkPeer, found := net.discoveredPeers[peerLink.Account]
			if !found {
				return fmt.Errorf("couldn't find %q in list of peers", peerLink.Account)
			}

			fmt.Println("adding weight to", peerLink.Account)
			// Weight defaults to 0.0
			peerLinkPeer.TotalWeight += peerLink.Weight
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

	net.orderedPeers = allPeers

	return nil
}

func (net *Network) OrderedPeers() []*Peer {
	return net.orderedPeers
}

func (net *Network) GetBlockHeight(height uint64) (blockhash string, err error) {
	resp, err := net.seedNetAPI.GetBlockByNum(height)
	if err != nil {
		return "", err
	}

	return resp.ID, nil
}

func (net *Network) PollGenesisTable() (data string, err error) {
	rows, err := net.seedNetAPI.GetTableRows(
		eos.GetTableRowsRequest{
			JSON:       true,
			Scope:      "eosio.disco",
			Code:       "eosio.disco",
			Table:      "genesis",
			TableKey:   "id",
			LowerBound: "",
			UpperBound: "",
			Limit:      1,
		},
	)
	if err != nil {
		return fmt.Errorf("get table rows: %s", err)
	}

	var candidates []*disco.Discovery
	if err := rows.JSONToStructs(&candidates); err != nil {
		return fmt.Errorf("reading discovery from table: %s", err)
	}

	for _, cand := range candidates {
		// verify cand discovery
		net.candidates[string(cand.SeedNetworkAccountName)] = cand
	}

}

func sha2(input []byte) string {
	hash := sha256.New()
	_, _ = hash.Write(input) // can't fail
	return hex.EncodeToString(hash.Sum(nil))
}

func (net *Network) PrintOrderedPeers() {
	fmt.Println("###############################################################################################")
	fmt.Println("####################################    PEER NETWORK    #######################################")
	fmt.Println("")

	columns := []string{
		"Role | IPNS Link | Account | Organization | Weight",
		"---- | --------- | ------- | ------------ | ------",
	}
	columns = append(columns, fmt.Sprintf("BIOS NODE | %s", net.orderedPeers[0].Columns()))
	for i := 1; i < 22 && len(net.orderedPeers) > i; i++ {
		columns = append(columns, fmt.Sprintf("ABP %02d | %s", i, net.orderedPeers[i].Columns()))
	}
	for i := 22; len(net.orderedPeers) > i; i++ {
		columns = append(columns, fmt.Sprintf("Part. %02d | %s", i, net.orderedPeers[i].Columns()))
	}
	fmt.Println(columnize.SimpleFormat(columns))

	fmt.Println("")
	fmt.Println("###############################################################################################")
	fmt.Println("")
}

// ReachedConsensus reads all the hashes of the top-level peers and
// returns true if we have reached an agreement on the content to
// inject in the chain.
func (net *Network) ReachedConsensus() bool {
	// TODO: Implement the logic that determines the consensus.. right
	// now it's just the weights in order.. and the top-most wins: we use
	// its configuration.
	return true
}

func (net *Network) ConsensusDiscovery() (*disco.Discovery, error) {
	// TODO: implement the algo to create a Discovery file based on
	// the most vouched for hashes for all the components.
	//
	// Will that work ? Will that make sense ?
	//
	// Cycle through the top peers, take the most vetted
	return net.orderedPeers[0].Discovery, nil
}
