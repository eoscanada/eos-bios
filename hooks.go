package bios

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/eoscanada/eos-bios/discovery"
)

var ConfiguredHooks = []HookDef{
	HookDef{"init", "Dispatch when we start the program."},
	HookDef{"boot_network", "Dispatched when we are BIOS Node, and our keys and node config is ready. Should trigger a config update and a restart."},
	HookDef{"publish_kickstart_data", "Dispatched with the contents of the (usually encrypted) Kickstart data, to be published to your social / web properties."},
	HookDef{"join_network", "Dispatched anyone joining the network. Could be as an Appointed Block Producer, or simply someone wanting to join the network after boot. It provides at least one p2p_address to connect to."},
	HookDef{"done", "When your process it done"},
}

type HookDef struct {
	Key  string
	Desc string
}

func (b *BIOS) DispatchInit() error {
	return b.dispatch("init", []string{}, nil)
}

func (b *BIOS) DispatchBootNetwork(genesisJSON, publicKey, privateKey string) error {
	return b.dispatch("boot_network", []string{
		"genesis_json", genesisJSON,
		"public_key", publicKey,
		"private_key", privateKey,
	}, nil)
}

func (b *BIOS) DispatchJoinNetwork(kickstart *KickstartData, peerDefs []*discovery.Peer) error {
	var names []string
	for _, peer := range peerDefs {
		names = append(names, peer.AccountName())
	}

	// Eventually, we might want to join a network and mesh directly with a few peers.
	// The hook won't have to change then..
	peerAddresses := []string{kickstart.BIOSP2PAddress}

	return b.dispatch("join_network", []string{
		"genesis_json", kickstart.GenesisJSON,
		"p2p_address_statements", "p2p-peer-address = " + strings.Join(peerAddresses, "\np2p-peer-address = "),
		"p2p_addresses", strings.Join(peerAddresses, ","),
		"producer_name_statements", "producer-name = " + strings.Join(names, "\nproducer-name = "),
		"producer_names", strings.Join(names, ","),
	}, nil)
}

func (b *BIOS) DispatchPublishKickstartData(kickstartData string) error {
	return b.dispatch("publish_kickstart_data", []string{
		"data", kickstartData,
	}, nil)
}

func (b *BIOS) DispatchDone() error {
	return b.dispatch("done", []string{}, nil)
}

// dispatch to both exec calls, and remote web hooks.
func (b *BIOS) dispatch(hookName string, data []string, f func() error) error {
	fmt.Printf("---- BEGIN HOOK %q ----\n", hookName)

	if len(data)%2 != 0 {
		return fmt.Errorf("data should be pairs of key and values, cannot have %d elements", len(data))
	}

	// check if `hook_[hookName]` exists or `hook_[hookName].sh` exists, and use that as a command,
	// otherwise, print that the hook is not present.
	filePaths := []string{
		fmt.Sprintf("./hook_%s", hookName),
		fmt.Sprintf("./hook_%s.sh", hookName),
	}
	var executable string
	for _, fl := range filePaths {
		if _, err := os.Stat(fl); err == nil {
			executable = fl
		}
	}

	if executable == "" {
		fmt.Printf("  - Hook not found (searched %q)\n", filePaths)
		return nil
	}

	args := []string{}
	for i := 0; i < len(data); i += 2 {
		v := data[i+1]
		args = append(args, v)
	}

	cmd := exec.Command(executable, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()

	//fmt.Printf("  Executing hook: %q\n", cmd.Args)

	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Printf("---- END HOOK %q ----\n", hookName)

	return nil
}
