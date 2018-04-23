package bios

import (
	"fmt"
	"strings"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
)

type LaunchData struct {
	LaunchBitcoinBlockHeight    int               `json:"launch_btc_block_height"`
	OpeningBalancesSnapshotHash string            `json:"opening_balances_snapshot_hash"`
	ContractHashes              map[string]string `json:"contract_hashes"`

	BootSequence []*OperationType `json:"boot_sequence"`

	Producers []*ProducerDef `json:"producers"`
}

type ProducerDef struct {
	// AccountName is the account we want to have created on the blockchain by the BIOS Boot node.
	AccountName eos.AccountName `json:"account_name"`

	// Authority is the original authority the Boot node will register
	// on that account. This allows teams to do their key ceremony a
	// few days before, and avoids a bootstrapping issue if we only
	// had a single public key for that account.
	Authority struct {
		Owner  eos.Authority `json:"owner"`
		Active eos.Authority `json:"active"`
	} `json:"authority"`

	// The key initially injected and used by the Appointed Block
	// Producers (if elected as such) to sign some of the first
	// blocks.
	//
	// When the ABP jumps in, it will `regproducer` with the same or a
	// different key (see Config's BlockSigningPublicKey).
	InitialBlockSigningPublicKey ecc.PublicKey `json:"initial_block_signing_key"`

	// KeybaseUser and PGPPublicKey are used to encrypt the Kickstart
	// Data payload, for the ABPs and followers.
	KeybaseUser  string `json:"keybase_user"`
	PGPPublicKey string `json:"pgp_public_key"`

	// OrganizationName is the block producer's name in plain text.
	OrganizationName string `json:"organization_name"`

	// Timezone, from https://en.wikipedia.org/wiki/List_of_tz_database_time_zones (column TZ)
	Timezone string `json:"timezone"`

	// Candidate producers are better off specifying a few URLs and social media properties, to avoid a single point of failure if they need to communicate with the world.
	URLs []string `json:"urls"`

	clonedFrom eos.AccountName
}

func (p *ProducerDef) String() string {
	return fmt.Sprintf("Account: % 15s   Keybase: % 32s   Org: % 30s   URLs: %s", p.AccountName, fmt.Sprintf("https://keybase.io/%s", p.KeybaseUser), p.OrganizationName, strings.Join(p.URLs, ", "))
}
