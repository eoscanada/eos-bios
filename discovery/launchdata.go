package discovery

type IPFSRef string
type IPNSRef string

type LaunchData struct {
	Peers                []PeerLink              `json:"peers"`
	BootSequence         IPFSRef                 `json:"boot_sequence"`
	Snapshot             IPFSRef                 `json:"snapshot"`
	SnapshotUnregistered IPFSRef                 `json:"snapshot_unregistered"`
	Contracts            map[string]ContractHash `json:"contracts"`
}

type PeerLink struct {
	DiscoveryFile IPNSRef `json:"discovery_url"`
	Comment       string  `json:"comment"`
	Weight        float64 `json:"weight"` // From 0.0 to 1.0
}

type ContractHash struct {
	ABI  IPFSRef `json:"abi"`
	Code IPFSRef `json:"code"`
}
