package bios

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/eoscanada/eos-bios/bios/disco"
	"github.com/eoscanada/eos-go"
	"github.com/ryanuber/columnize"
	"github.com/spf13/viper"
)

type Network struct {
	SingleOnly bool

	MyPeer *Peer

	SeedNetAPI      *eos.API
	seedNetContract string

	ipfs            *IPFS
	cachePath       string
	discoveredPeers map[eos.AccountName]*Peer
	orderedPeers    []*Peer
	candidates      map[string]*disco.Discovery
}

func NewNetwork(cachePath string, myDiscovery *disco.Discovery, ipfs *IPFS, seedNetContract string, seedNetAPI *eos.API) *Network {
	return &Network{
		SeedNetAPI:      seedNetAPI,
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

	var rows []struct {
		ID            eos.AccountName  `json:"id"`
		DiscoveryFile *disco.Discovery `json:"content"`
		UpdatedAt     eos.JSONTime     `json:"updated_at"`
	}

	if !net.SingleOnly {
		fmt.Println("Updating network graph")
		rowsJSON, err := net.SeedNetAPI.GetTableRows(
			eos.GetTableRowsRequest{
				JSON:     true,
				Scope:    net.seedNetContract,
				Code:     net.seedNetContract,
				Table:    "discovery",
				TableKey: "id",
				//LowerBound: "",
				//UpperBound: "",
				Limit: 1000,
			},
		)
		if err != nil {
			return fmt.Errorf("get discovery rows: %s", err)
		}

		if err := rowsJSON.JSONToStructs(&rows); err != nil {
			return fmt.Errorf("reading discovery from table: %s", err)
		}
	}

	for _, cand := range rows {
		// TODO: verify their discovery file.. the values in there.. do we simply skip those with invalid weights for example ?? They're excluded from the graph?

		cand.DiscoveryFile.UpdatedAt = cand.UpdatedAt
		cand.DiscoveryFile.SeedNetworkAccountName = cand.ID // we override what they think, we use what they *signed* for..
		net.candidates[string(cand.ID)] = cand.DiscoveryFile
	}

	if err := net.ensureExists(); err != nil {
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

	if viper.GetBool("verbose") {
		fmt.Printf("- %q has %d peer(s)\n", discoFile.SeedNetworkAccountName, len(discoFile.SeedNetworkPeers))
	}

	for _, peerLink := range discoFile.SeedNetworkPeers {
		if viper.GetBool("verbose") {
			fmt.Printf("  - peer %s comment=%q, weight=%d\n", peerLink.Account, peerLink.Comment, peerLink.Weight)
		}

		peerDisco, found := net.candidates[string(peerLink.Account)]
		if !found {
			if viper.GetBool("verbose") {
				fmt.Println("    - peer not found, won't weight in")
			}
			continue
		}

		if net.discoveredPeers[peerDisco.SeedNetworkAccountName] != nil {
			if viper.GetBool("verbose") {
				fmt.Printf("    - already added %q\n", peerDisco.SeedNetworkAccountName)
			}
			continue
		}
		if viper.GetBool("verbose") {
			fmt.Printf("    - adding %q\n", peerDisco.SeedNetworkAccountName)
		}

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

	fmt.Printf("Downloading and caching content from IPFS: %q\n", ref)
	cnt, err := net.ipfs.Get(ref)
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

func (net *Network) ReaderFromCache(ref string) (io.ReadCloser, error) {
	fileName := replaceAllWeirdities(ref)
	return os.Open(filepath.Join(net.cachePath, fileName))
}

func (net *Network) FileNameFromCache(ref string) string {
	fileName := replaceAllWeirdities(ref)
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

		if viper.GetBool("verbose") {
			fmt.Println("First level peer", peer.AccountName())
		}
		for _, peerLink := range peer.Discovery.SeedNetworkPeers {

			if peer.Discovery.SeedNetworkAccountName == peerLink.Account {
				// Can't vote for yourself
				continue
			}

			peerLinkPeer, found := net.discoveredPeers[peerLink.Account]
			if !found {
				continue
			}

			if peerLink.Weight <= 100 {
				peerLinkPeer.TotalWeight += int(peerLink.Weight)
			}
		}

		allPeers = append(allPeers, peer)
	}

	// Sort the `orderedPeers`
	sort.Slice(allPeers, func(i, j int) bool {
		if allPeers[i].TotalWeight == allPeers[j].TotalWeight {
			return allPeers[i].AccountName() < allPeers[j].AccountName()
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
	resp, err := net.SeedNetAPI.GetBlockByNum(height)
	if err != nil {
		return "", err
	}

	return resp.ID, nil
}

func (net *Network) PollGenesisTable(account eos.AccountName) (data string, err error) {
	accountRaw, err := eos.MarshalBinary(account)
	if err != nil {
		return "", err
	}
	accountInt := binary.LittleEndian.Uint64(accountRaw)
	rowsJSON, err := net.SeedNetAPI.GetTableRows(
		eos.GetTableRowsRequest{
			JSON:       true,
			Scope:      net.seedNetContract,
			Code:       net.seedNetContract,
			Table:      "genesis",
			TableKey:   "id",
			LowerBound: fmt.Sprintf("%d", accountInt),
			UpperBound: fmt.Sprintf("%d", accountInt+1),
			Limit:      1,
		},
	)
	if err != nil {
		return "", fmt.Errorf("get genesis rows: %s", err)
	}

	var rows []struct {
		ID                  eos.AccountName `json:"id"`
		GenesisJSON         string          `json:"genesis_json"`
		InitialP2PAddresses []string        `json:"initial_p2p_addresses"`
		UpdatedAt           eos.JSONTime    `json:"updated_at"`
	}
	if err := rowsJSON.JSONToStructs(&rows); err != nil {
		return "", fmt.Errorf("reading discovery from table: %s", err)
	}

	if len(rows) != 1 {
		return "", nil
	}

	return rows[0].GenesisJSON, nil
}

func (net *Network) PrintDiscoveryFiles() (err error) {
	fmt.Println("List of all accounts that have published a discovery file:")
	for _, cand := range net.candidates {
		fmt.Printf("%s\n", cand.SeedNetworkAccountName)
	}
	return
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
		"Role | Seed Account | Target Acct | Weight | Offset | Block Height",
		"---- | ------------ | ----------- | ------ | ------ | ------------",
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
