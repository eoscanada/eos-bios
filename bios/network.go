package bios

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/eoscanada/eos-bios/bios/disco"
	"github.com/eoscanada/eos-go"
	"github.com/ryanuber/columnize"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

type Network struct {
	Log *Logger

	MyPeer *Peer

	SeedNetAPI      *eos.API
	seedNetContract string

	// nodes format
	allNodes          *simple.WeightedDirectedGraph
	allNodesFetchFunc func() error
	allNetworks       []*simple.WeightedDirectedGraph
	myNetwork         *simple.WeightedDirectedGraph

	ipfs           *IPFS
	ipfsReferences []ipfsRef
	cachePath      string
}

type ipfsRef struct {
	Name          string
	Reference     string
	SourceAccount string
}

func NewNetwork(cachePath string, myDiscovery *disco.Discovery, ipfs *IPFS, seedNetContract string, seedNetAPI *eos.API) *Network {
	net := &Network{
		SeedNetAPI:      seedNetAPI,
		ipfs:            ipfs,
		cachePath:       cachePath,
		seedNetContract: seedNetContract,
		MyPeer: &Peer{
			Discovery: myDiscovery,
			UpdatedAt: time.Now(),
		},
	}
	net.allNodesFetchFunc = net.fetchGraphFromSeedNetwork
	return net
}

func (net *Network) SetLocalNetwork() {
	net.allNodesFetchFunc = net.fetchSingleNode
}

func (net *Network) UpdateGraph() error {
	net.allNodes = simple.NewWeightedDirectedGraph(0, 0)

	if err := net.allNodesFetchFunc(); err != nil {
		return err
	}

	for _, node := range net.allNodes.Nodes() {
		peer := node.(*Peer)
		if err := net.loadTargetContentsRefs(peer); err != nil {
			return fmt.Errorf("loading target contents refs: %s", err)
		}

		if err := net.traversePeers(peer); err != nil {
			return fmt.Errorf("traversing peers: %s", err)
		}
	}

	net.isolateNetworks()

	net.calculateNetworkWeights()

	return nil
}

func (net *Network) fetchSingleNode() error {
	newPeer := &Peer{
		UpdatedAt: time.Now(),
		Discovery: net.MyPeer.Discovery,
	}
	net.allNodes.AddNode(newPeer)

	return nil
}

func (net *Network) fetchGraphFromSeedNetwork() error {
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

	var rows []struct {
		ID        eos.AccountName  `json:"id"`
		Discovery *disco.Discovery `json:"content"`
		UpdatedAt eos.JSONTime     `json:"updated_at"`
	}
	if err := rowsJSON.JSONToStructs(&rows); err != nil {
		return fmt.Errorf("reading discovery from table: %s", err)
	}

	for _, cand := range rows {
		if err := ValidateDiscovery(cand.Discovery); err != nil {
			fmt.Printf("Skipping invalid discovery file from %q: %s\n", cand.ID, err)
			continue
		}

		// cand.Discovery.UpdatedAt = cand.UpdatedAt
		cand.Discovery.SeedNetworkAccountName = cand.ID // we override what they think, we use what they *signed* for..

		newPeer := &Peer{
			UpdatedAt: cand.UpdatedAt.Time,
			Discovery: cand.Discovery,
		}
		if !net.allNodes.Has(newPeer.ID()) { // rows can't have duplicate key anyway
			net.allNodes.AddNode(newPeer)
		}
	}

	return nil
}

func (net *Network) loadTargetContentsRefs(peer *Peer) error {
	for _, contentRef := range peer.Discovery.TargetContents {
		if contentRef.Ref == "" {
			net.Log.Debugf("  - WARN: %q has an empty ipfs ref for name=%q\n", peer.Discovery.SeedNetworkAccountName, contentRef.Name)
			continue
		}

		if !strings.HasPrefix(contentRef.Ref, "/ipfs/") {
			net.Log.Debugf("  - WARN: %q has a ref that doesn't start with '/ipfs/' for name=%q\n", peer.Discovery.SeedNetworkAccountName, contentRef.Name)
			continue
		}

		net.ipfsReferences = append(net.ipfsReferences, ipfsRef{
			Name:          contentRef.Name,
			Reference:     contentRef.Ref,
			SourceAccount: string(peer.Discovery.SeedNetworkAccountName),
		})
	}

	return nil
}

func (net *Network) traversePeers(fromPeer *Peer) error {
	net.Log.Debugf("- %q has %d peer(s)\n", fromPeer.Discovery.SeedNetworkAccountName, len(fromPeer.Discovery.SeedNetworkPeers))

	for _, peerLink := range fromPeer.Discovery.SeedNetworkPeers {
		net.Log.Debugf("  - peer %s comment=%q, weight=%d\n", peerLink.Account, peerLink.Comment, peerLink.Weight)

		peerID := AccountToNodeID(peerLink.Account)
		if !net.allNodes.Has(peerID) {
			net.Log.Debugln("    - peer not found, won't weight in")
			continue
		}

		toPeer := net.allNodes.Node(peerID).(*Peer)
		peerEdge := &PeerEdge{
			FromPeer: fromPeer,
			ToPeer:   toPeer,
			PeerLink: peerLink,
		}

		if fromPeer == toPeer {
			net.Log.Debugf("    - no self-ref allowed\n")
			continue
		}

		if net.allNodes.HasEdgeFromTo(fromPeer.ID(), toPeer.ID()) {
			net.Log.Debugf("    - duplicate link to %q\n", toPeer.Discovery.SeedNetworkAccountName)
			continue
		}

		net.Log.Debugf("    - adding %q\n", toPeer.Discovery.SeedNetworkAccountName)

		net.allNodes.SetWeightedEdge(peerEdge)
	}

	return nil
}

func (net *Network) isolateNetworks() {
	net.allNetworks = make([]*simple.WeightedDirectedGraph, 0, 0)
	// myAccount := net.MyPeer.Discovery.SeedNetworkAccountName

	for _, subnet := range topo.TarjanSCC(net.allNodes) {
		subGraph := simple.NewWeightedDirectedGraph(0, 0)

		// Grab the nodes
		for _, node := range subnet {
			subGraph.AddNode(node)
		}

		// Grab only the edges that fit the subgraph
		for _, edge := range net.allNodes.WeightedEdges() {
			if subGraph.Has(edge.From().ID()) && subGraph.Has(edge.To().ID()) {
				subGraph.SetWeightedEdge(edge)
			}
		}

		net.allNetworks = append(net.allNetworks, subGraph)
	}
}

//
// Assets download and caching
//

func (net *Network) DownloadReferences() error {
	if err := net.ensureCacheExists(); err != nil {
		return fmt.Errorf("error creating cache path: %s", err)
	}

	for _, contentRef := range net.ipfsReferences {
		if err := net.DownloadIPFSRef(contentRef.Reference); err != nil {
			return fmt.Errorf("content %q: %s", contentRef.Name, err)
		}
	}
	return nil
}

func (net *Network) ensureCacheExists() error {
	return os.MkdirAll(net.cachePath, 0777)
}

func (net *Network) DownloadIPFSRef(ref string) error {
	if net.isInCache(string(ref)) {
		//fmt.Printf("ipfs ref: %q in cache\n", ref)
		return nil
	}

	fmt.Printf("Downloading and caching content from IPFS: %q\n", ref)
	cnt, err := net.ipfs.Get(ref)
	if err != nil {
		return err
	}

	if err := net.writeToCache(ref, cnt); err != nil {
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

//
// Graph weighting...
//

func (net *Network) calculateNetworkWeights() {
	// For all networks
	for _, network := range net.allNetworks {

		for _, node := range network.Nodes() {
			var totalWeight int
			for _, inwardNode := range network.To(node.ID()) {
				edge := network.WeightedEdge(inwardNode.ID(), node.ID())
				totalWeight += int(edge.Weight())
			}
			node.(*Peer).TotalWeight = totalWeight
		}
	}
}

func (net *Network) NetworkThatIncludes(networkAccount eos.AccountName) *simple.WeightedDirectedGraph {
	for _, network := range net.allNetworks {
		if !network.Has(AccountToNodeID(networkAccount)) {
			continue // not my network Jack !
		}
		return network
	}

	return nil
}

func (net *Network) OrderedPeers(network *simple.WeightedDirectedGraph) (out []*Peer) {
	if network == nil {
		return
	}

	for _, node := range network.Nodes() {
		out = append(out, node.(*Peer))
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i].TotalWeight == out[j].TotalWeight {
			return out[i].AccountName() < out[j].AccountName()
		}
		return out[i].TotalWeight > out[j].TotalWeight
	})

	return
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

func sha2(input []byte) string {
	hash := sha256.New()
	_, _ = hash.Write(input) // can't fail
	return hex.EncodeToString(hash.Sum(nil))
}

func (net *Network) ListNetworks(verbose bool) {
	fmt.Println("Networks formed by published discovery files:")

	for idx, network := range net.allNetworks {
		fmt.Printf("%d.\n", idx+1)
		orderedPeers := net.OrderedPeers(network)
		for _, peer := range orderedPeers {
			fmt.Printf("  - %s (total weight: %d)\n", peer.Discovery.SeedNetworkAccountName, peer.TotalWeight)
		}
	}
}

func (net *Network) MyNetwork() *simple.WeightedDirectedGraph {
	network := net.NetworkThatIncludes(net.MyPeer.Discovery.SeedNetworkAccountName)
	if network == nil {
		if len(net.MyPeer.Discovery.SeedNetworkPeers) == 0 {
			fmt.Println("You are part of no network. Either define a `seed_network_peers` to point to some peers in a network, or ask to be pointed to by someone in a network")
			os.Exit(1)
		}

		network = net.NetworkThatIncludes(net.MyPeer.Discovery.SeedNetworkPeers[0].Account)
		if network == nil {
			fmt.Println("You're part of no network, and your first peer in `seed_network_peers` isn't either (!!)")
			os.Exit(1)
		}
	}

	return network
}

func (net *Network) PrintOrderedPeers() {
	fmt.Println("###############################################################################################")
	fmt.Println("####################################    PEER NETWORK    #######################################")
	fmt.Println("")

	network := net.MyNetwork()
	orderedPeers := net.OrderedPeers(network)

	contentAgreement := ComputeContentsAgreement(orderedPeers)
	peerContent := ComputePeerContentsColumn(contentAgreement, orderedPeers)

	columns := []string{
		"Role | Seed Account | Target Acct | Weight | GMT | Launch Block | Contents",
		"---- | ------------ | ----------- | ------ | --- | ------------ | --------",
	}
	columns = append(columns, fmt.Sprintf("BIOS NODE | %s | %s", orderedPeers[0].Columns(), peerContent[0]))
	for i := 1; i < 22 && len(orderedPeers) > i; i++ {
		columns = append(columns, fmt.Sprintf("ABP %02d | %s | %s", i, orderedPeers[i].Columns(), peerContent[i]))
	}
	for i := 22; len(orderedPeers) > i; i++ {
		columns = append(columns, fmt.Sprintf("Part. %02d | %s | %s", i, orderedPeers[i].Columns(), peerContent[i]))
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

	orderedPeers := net.OrderedPeers(net.MyNetwork())
	return orderedPeers[0].Discovery, nil
}
