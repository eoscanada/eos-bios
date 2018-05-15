package disco

import (
	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
)

// PeerLink is the struct specified by the user
type PeerLink struct {
	Account eos.AccountName `json:"account"`
	Comment string          `json:"comment"`
	Weight  uint8           `json:"weight"`
}

type Discovery struct {
	SeedNetworkChainID                     eos.SHA256Bytes `json:"seed_network_chain_id"`
	SeedNetworkAccountName                 eos.AccountName `json:"seed_network_account_name"`
	SeedNetworkHTTPAddress                 string          `json:"seed_network_http_address"`
	SeedNetworkPeers                       []*PeerLink     `json:"seed_network_peers"`
	SeedNetworkLaunchBlock                 uint64          `json:"seed_network_launch_block"`
	URLs                                   []string        `json:"urls"`
	GMTOffset                              int16           `json:"gmt_offset"`
	TargetNetworkIsTest                    uint8           `json:"target_network_is_test"`
	TargetChainID                          eos.SHA256Bytes `json:"target_chain_id"`
	TargetP2PAddress                       string          `json:"target_p2p_address"`
	TargetHTTPAddress                      string          `json:"target_http_address"`
	TargetAccountName                      eos.AccountName `json:"target_account_name"`
	TargetAppointedBlockProducerSigningKey ecc.PublicKey   `json:"target_appointed_block_producer_signing_key"`
	TargetInitialAuthority                 struct {
		Owner  eos.Authority `json:"owner"`
		Active eos.Authority `json:"active"`
	} `json:"target_initial_authority"`

	TargetContents []ContentRef `json:"target_contents"`
}

type ContentRef struct {
	Name    string `json:"name"`
	Ref     string `json:"ref"`
	Comment string `json:"comment"`
}

// DiscoveryRow represents a row in the `eosio.disco` contract, for the `discovery` table.
type DiscoveryRow struct {
	ID        eos.AccountName `json:"id"`
	Content   *Discovery      `json:"content"`
	UpdatedAt eos.JSONTime    `json:"updated_at"`
}

// GenesisRow represents a row in the `eosio.disco` contract, for the `genesis` table.
type GenesisRow struct {
	ID                  eos.AccountName `json:"id"`
	GenesisJSON         string          `json:"genesis_json"`
	InitialP2PAddresses []string        `json:"initial_p2p_addresses"`
}
