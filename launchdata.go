package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type LaunchData struct {
	LaunchBitcoinBlockHeight    int             `json:"launch_btc_block_height" yaml:"launch_btc_block_height"`
	OpeningBalancesSnapshotHash string          `json:"opening_balances_snapshot_hash" yaml:"opening_balances_snapshot_hash"`
	SystemContractHash          string          `json:"system_contract_hash" yaml:"system_contract_hash"`
	Producers                   []*ProducerData `json:"producers" yaml:"producers"`
}
type ProducerData struct {
	EOSIOAccountName string `json:"eosio_account_name" yaml:"eosio_account_name"`
	EOSIOPublicKey   string `json:"eosio_public_key" yaml:"eosio_public_key"`
	KeybaseUser      string `json:"keybase_user" yaml:"keybase_user"`
	PGPPublicKey     string `json:"pgp_public_key" yaml:"pgp_public_key"`
	AgentName        string `json:"agent_name" yaml:"agent_name"`
	URL              string `json:"url" yaml:"url"`
}

// snapshotPath, codePath, abiPath string
func loadLaunchFile(filename string, config *Config) (out *LaunchData, err error) {
	cnt, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(cnt, &out); err != nil {
		return nil, err
	}

	if out.LaunchBitcoinBlockHeight == 0 {
		return nil, fmt.Errorf("launch_btc_block_height unspecified (or 0)")
	}

	// Hash the `--opening-balance-snapshot` file, compare to `launch.
	snapshotHash, err := hashFile(config.OpeningBalances.SnapshotPath)
	if err != nil {
		return nil, err
	}

	log.Printf("Hash of %q: %s", config.OpeningBalances.SnapshotPath, snapshotHash)

	if snapshotHash != out.OpeningBalancesSnapshotHash {
		return nil, fmt.Errorf("snapshot hash doesn't match launch data")
	}

	codeHash, err := hashCodeFiles(config.SystemContract.CodePath, config.SystemContract.ABIPath)
	if err != nil {
		return nil, fmt.Errorf("error hashing system contract's code + abi: %s", err)
	}

	log.Printf("Hash of %q and %q: %s", config.SystemContract.CodePath, config.SystemContract.ABIPath, codeHash)

	if codeHash != out.SystemContractHash {
		return nil, fmt.Errorf("system contract's code hash don't match")
	}

	// Verify the `producers` entries's public keys start with `EOS`
	// and are the right length, etc.. try to load them.

	// Check duplicate entries in `launch.yaml`, fail immediately.
	//    Check the `eosio_account_name`
	// Hash the eosio-system-code and eosio-system-abi files, concatenated.
	//    If check fails, print the hash.. always print the hash.

	return out, nil
}

func hashFile(filename string) (string, error) {
	h := sha256.New()

	cnt, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	h.Write(cnt)

	return hex.EncodeToString(h.Sum(nil)), nil
}

func hashCodeFiles(code, abi string) (string, error) {
	h := sha256.New()

	cnt, err := ioutil.ReadFile(code)
	if err != nil {
		return "", err
	}

	h.Write(cnt)

	h.Write([]byte(":"))

	cnt, err = ioutil.ReadFile(abi)
	if err != nil {
		return "", err
	}

	h.Write(cnt)

	return hex.EncodeToString(h.Sum(nil)), nil
}
