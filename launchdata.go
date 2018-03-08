package main

type LaunchData struct {
	LaunchBitcoinBlockHeight    int             `json:"launch_btc_block_height"`
	OpeningBalancesSnapshotHash string          `json:"opening_balances_snapshot_hash"`
	SystemContracthash          string          `json:"system_contract_hash"`
	Producers                   []*ProducerData `json:"producers"`
}
type ProducerData struct {
	EOSIOAccountName string `json:"eosio_account_name"`
	EOSIOPublicKey   string `json:"eosio_public_key"`
	KeybaseUser      string `json:"keybase_user"`
	PGPPublicKey     string `json:"pgp_public_key"`
	AgentName        string `json:"agent_name"`
	URL              string `json:"url"`
}
