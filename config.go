package bios

import (
	"net/url"

	"github.com/eoscanada/eos-go/ecc"
	"github.com/eoscanada/eos-go/system"
)

type Config struct {
	// Producer describes your producing node.
	Peer struct {
		// MyAccount is the name of the `account_name` this producer will be using on chain
		MyAccount string `json:"my_account"`
		// APIAddress is the target API endpoint for the locally booting node, a clean-slate node. It can be routable only from the local machine.
		APIAddress    string   `json:"api_address"`
		APIAddressURL *url.URL `json:"-"`
		// SecretP2PAddress is the endpoint which will be published at the end of the process. Needs to be externally routable.  It must be kept secret for DDoS protection.
		SecretP2PAddress string `json:"secret_p2p_address"`

		// Key you want to register to sign blocks.
		BlockSigningPublicKey ecc.PublicKey `json:"block_signing_public_key"`

		// Private key used to setup your config.ini as an ABP (or as someone joining the network)
		BlockSigningPrivateKeyPath string `json:"block_signing_private_key_path"`

		// Available once loaded successfuly from the previous field's path.
		BlockSigningPrivateKey *ecc.PrivateKey `json:"-"`
	} `json:"producer"`

	// Hooks are called at different stages in the process, for
	// remote systems to be notified and act.  They are simply `http`
	// endpoints to which a POST will be sent with pre-defined structs
	// as JSON.  See `hooks.go`
	Hooks map[string]*HookConfig `json:"hooks"`

	MyParameters system.EOSIOParameters `json:"my_parameters"`

	// PGP manages the PGP keys, used for the communications channel.
	PGP struct {
		// Program represents the type of program to use (gpg, keybase ?)
		Program string `json:"program"`
		// Path to binary executable.. unless we use in-process cryptography..
		Path string `json:"path"`
	} `json:"pgp"`
}

type HookConfig struct {
	URL  string `json:"url"`
	Exec string `json:"exec"`
	Wait bool   `json:"wait"`
}

/*

Default values in code for eosio_parameters:

static const uint32_t bandwidth_average_window_ms   = 24*60*60*1000l;
static const uint32_t compute_average_window_ms     = 24*60*60*1000l;
static const uint32_t blocksize_average_window_ms   = 60*1000l;


const static uint32_t   default_max_block_size              = 1024 * 1024; /// at 500ms blocks and 200byte trx, this enables 10,000 TPS burst
const static uint32_t   default_target_block_size           = default_max_block_size / 10; /// we target 1000 TPS burst
const static uint32_t   default_target_block_acts_per_scope  = 1000;
const static uint32_t   default_max_block_acts_per_scope     = default_target_block_acts_per_scope*10;

const static uint32_t   default_target_block_acts  = 2000;
const static uint32_t   default_max_block_acts     = default_target_block_acts*100;
const static uint32_t   setcode_act_usage          = 100;

const static uint64_t   default_max_storage_size       = 10 * 1024;
const static uint32_t   default_max_trx_lifetime       = 60*60;
const static uint16_t   default_max_auth_depth         = 6;
const static uint32_t   default_max_trx_runtime        = 10*1000;
const static uint16_t   default_max_inline_depth       = 4;
const static uint32_t   default_max_inline_action_size = 4 * 1024;
const static uint32_t   default_max_gen_trx_size       = 64 * 1024; ///
const static uint32_t   default_max_gen_trx_count      = 16; ///< the number of generated transactions per action
const static uint32_t   producers_authority_threshold  = 14;
const static uint32_t   rate_limiting_precision        = 1000*1000;

const static uint16_t   max_recursion_depth = 6;
*/
