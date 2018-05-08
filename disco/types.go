package disco

import (
	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
)

type PeerLink struct {
	Account eos.AccountName `json:"account"`
	Comment string          `json:"comment"`
	Weight  float64         `json:"weight"`
}
type ContentRef struct {
	Name    string `json:"name"`
	Ref     string `json:"ref"`
	Comment string `json:"comment"`
}

type Discovery struct {
	SeedNetworkChainID                     eos.SHA256Bytes `json:"seed_network_chain_id"`
	SeedNetworkAccountName                 eos.AccountName `json:"seed_network_account_name"`
	SeedNetworkPeers                       []*PeerLink     `json:"seed_network_peers"`
	SeedNetworkLaunchBlock                 uint64          `json:"seed_network_launch_block"`
	URLs                                   []string        `json:"urls"`
	GMTOffset                              int16           `json:"gmt_offset"`
	TargetNetworkIsTest                    int8            `json:"target_network_is_test"`
	TargetChainID                          eos.SHA256Bytes `json:"target_chain_id"`
	TargetP2PAddress                       string          `json:"target_p2p_address"`
	TargetAccountName                      eos.AccountName `json:"target_account_name"`
	TargetAppointedBlockProducerSigningKey ecc.PublicKey   `json:"target_appointed_block_producer_signing_key"`
	TargetInitialAuthority                 struct {
		Owner    eos.Authority `json:"owner"`
		Active   eos.Authority `json:"active"`
		Recovery eos.Authority `json:"recovery"`
	} `json:"target_initial_authority"`

	TargetContents []ContentRef `json:"target_contents"`

	UpdatedAt eos.JSONTime `json:"-"` // injected in `UpdatedGraph`
}

type GenesisRow struct {
	ID                  eos.AccountName `json:"id"`
	GenesisJSON         string          `json:"genesis_json"`
	InitialP2PAddresses []string        `json:"initial_p2p_addresses"`
	UpdatedAt           eos.Tstamp      `json:"updated_at"` // or JSONTime in 32 bits ?
}

type DiscoveryRow struct {
	ID        eos.AccountName `json:"id"`
	Content   Discovery       `json:"content"`
	UpdatedAt eos.Tstamp      `json:"updated_at"`
}
