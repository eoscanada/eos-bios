package main

import (
	"fmt"
	"io/ioutil"

	"text/template"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	SystemContract struct {
		CodePath string `json:"code_path" yaml:"code_path"`
		ABIPath  string `json:"abi_path" yaml:"abi_path"`
		Hash     string `json:"hash" yaml:"hash"`
	} `json:"system_contract" yaml:"system_contract"`

	// Producer describes your producing node.
	Producer struct {
		// MyAccount is the name of the `account_name` this producer will be using on chain
		MyAccount string `json:"my_account" yaml:"my_account"`
		// APIAddress is the target API endpoint for the locally booting node, a clean-slate node. It can be routable only from the local machine.
		APIAddress string `json:"api_address" yaml:"api_address"`
		// SecretP2PAddress is the endpoint which will be published at the end of the process. Needs to be externally routable.  It must be kept secret for DDoS protection.
		SecretP2PAddress string `json:"secret_p2p_address" yaml:"secret_p2p_address"`

		// WalletAddress is the API endpoint where your wallet lives
		WalletAddress string `json:"wallet_address" yaml:"wallet_address"`
	} `json:"producer" yaml:"producer"`

	// OpeningBalancesSnapshotPath represents the `snapshot.csv` file,
	// which holds the opening balances for all ERC-20 token holders.
	OpeningBalances struct {
		// SnapshotPath is the path to the `csv` file, extracted using the `genesis` tool.
		SnapshotPath string `json:"snapshot_path" yaml:"snapshot_path"`
	} `json:"opening_balances" yaml:"opening_balances"`

	// PGP manages the PGP keys, used for the communications channel.
	PGP struct {
		// Whether to use Keybase, or simply use in-built PGP crypto.
		UseKeybase bool `json:"use_keybase" yaml:"use_keybase"`
		// If `UseKeybase` is false, provide your secret `KeyPath` here.
		KeyPath string `json:"key_path" yaml:"key_path"`
	} `json:"pgp" yaml:"pgp"`

	NoShuffle bool `json:"disable_shuffling" yaml:"disable_shuffling"`

	// Hooks are called at different stages in the process, for
	// remote systems to be notified and act.  They are simply `http`
	// endpoints to which a POST will be sent with pre-defined structs
	// as JSON.  See `hooks.go`
	Hooks map[string]*HookConfig `json:"connect_to_bios"`
}

type HookConfig struct {
	URL          string `json:"url"`
	Exec         string `json:"exec"`
	execTemplate *template.Template
	TweetAPIKey  string `json:"tweet_api_key"` // for example, we could have those things support tweets !
}

func (c HookConfig) parseTemplate() (*template.Template, error) {
	if c.Exec == "" {
		return nil, nil
	}

	tpl, err := template.New("exec").Parse(c.Exec)
	return tpl, err
}

func LoadLocalConfig(localConfigPath string) (*Config, error) {
	cnt, err := ioutil.ReadFile(localConfigPath)
	if err != nil {
		return nil, err
	}

	var c *Config
	if err = yaml.Unmarshal(cnt, &c); err != nil {
		return nil, err
	}

	// TODO: do more checks on configuration...
	// TODO: test all Webhook URLs if defined
	// TODO: test all Hooks's Exec templates, and compile them right away..
	h := c.Hooks
	fmt.Println("Hooks runtime config (see `hooks.go`):")
	for _, hook := range configuredHooks {
		hconf := h[hook.Key]
		if hconf == nil {
			fmt.Printf("Hook %q NOT configured\n", hook.Key)
		}

		if hconf.Exec != "" {
			fmt.Printf("Hook %q configured to EXEC\n", hook.Key)
		}
		if hconf.URL != "" {
			fmt.Printf("Hook %q configured to POST via HTTP\n", hook.Key)
		}

		tpl, err := hconf.parseTemplate()
		if err != nil {
			return nil, fmt.Errorf(`Hook %q parsing FAILED: invalid "exec" template: %s`, hook.Key, err)
		}

		hconf.execTemplate = tpl
	}

	return c, nil
}
