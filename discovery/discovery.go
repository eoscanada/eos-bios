package discovery

type Discovery struct {
	// Testnet is true if this discovery file represents a testing
	// network.
	Testnet bool `json:"testnet"`
	// Mainnet is true if this discovery file represents the main net
	// (or a production network). One of Testnet and Mainnet must be
	// `true`, and are mutually exclusive.
	Mainnet bool `json:"mainnet"`

	EOSIOAccountName   string     `json:"eosio_account_name"`
	EOSIOABPSigningKey string     `json:"eosio_appointed_block_producer_signing_key"`
	LaunchData         LaunchData `json:"launch_data"`
}
