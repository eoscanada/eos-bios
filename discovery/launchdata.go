package discovery

type LaunchData struct {
	Wingmen              []Wingman               `json:"wingmen"`
	BootSequence         HashURL                 `json:"boot_sequence"`
	Snapshot             HashURL                 `json:"snapshot"`
	SnapshotUnregistered HashURL                 `json:"snapshot_unregistered"`
	Contracts            map[string]ContractHash `json:"contracts"`
}

type Wingman struct {
	DiscoveryURL string  `json:"discovery_url"`
	Comment      string  `json:"comment"`
	Weight       float64 `json:"weight"` // From 0.0 to 1.0
}

type HashURL struct {
	// Hash is the important bit, used to weight in votes from
	// different BPs.
	Hash string `json:"hash"`
	// URL is the downloadable location of the resource, which should
	// hash to `Hash`.
	URL string `json:"url"`
	// Comment has any relevant details, how the resource was built,
	// produced, who did it, when, based on what revision of which
	// source code, and whatnot.
	Comment string `json:"comment"`
}

type ContractHash struct {
	ABI  HashURL `json:"abi"`
	Code HashURL `json:"code"`
}
