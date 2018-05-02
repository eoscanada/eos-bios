package bios

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func (b *BIOS) DispatchInit(operation string) error {
	return b.dispatch("init", []string{
		operation,
	}, nil)
}

func (b *BIOS) DispatchBootNetwork(genesisJSON, publicKey, privateKey string) error {
	return b.dispatch("boot_network", []string{
		genesisJSON,
		publicKey,
		privateKey,
	}, nil)
}

func (b *BIOS) DispatchJoinNetwork(kickstart *KickstartData, peerDefs []*Peer, otherPeers []string) error {
	var names []string
	for _, peer := range peerDefs {
		names = append(names, peer.AccountName())
	}

	return b.dispatch("join_network", []string{
		kickstart.GenesisJSON,
		"p2p-peer-address = " + strings.Join(otherPeers, "\np2p-peer-address = "),
		strings.Join(otherPeers, ","),
		"producer-name = " + strings.Join(names, "\nproducer-name = "),
		strings.Join(names, ","),
	}, nil)
}

func (b *BIOS) DispatchPublishKickstartData(kickstartData string) error {
	return b.dispatch("publish_kickstart_data", []string{
		kickstartData,
	}, nil)
}

func (b *BIOS) DispatchDone() error {
	return b.dispatch("done", []string{}, nil)
}

// dispatch to both exec calls, and remote web hooks.
func (b *BIOS) dispatch(hookName string, args []string, f func() error) error {
	fmt.Printf("---- BEGIN HOOK %q ----\n", hookName)

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
