package bios

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func (b *BIOS) DispatchInit(operation string) error {
	return b.dispatch("init", []string{
		operation, // "join", "orchestrate", "boot"
	}, nil)
}

func (b *BIOS) DispatchBootPublishGenesis(genesisJSON string) error {
	encodedGenesis := base64.RawStdEncoding.EncodeToString([]byte(genesisJSON))

	return b.dispatch("boot_publish_genesis", []string{
		encodedGenesis,
		genesisJSON,
	}, nil)
}

func (b *BIOS) DispatchBootNode(genesisJSON, publicKey, privateKey string, otherPeers []string) error {
	return b.dispatch("boot_node", []string{
		genesisJSON,
		publicKey,
		privateKey,
		"p2p-peer-address = " + strings.Join(otherPeers, "\np2p-peer-address = "),
		strings.Join(otherPeers, ","),
	}, nil)
}

func (b *BIOS) DispatchJoinNetwork(genesis *GenesisJSON, peerDefs []*Peer, otherPeers []string) error {
	var names []string
	for _, peer := range peerDefs {
		names = append(names, string(peer.Discovery.TargetAccountName))
	}

	cnt, err := json.Marshal(genesis)
	if err != nil {
		return err
	}

	return b.dispatch("join_network", []string{
		string(cnt),
		"p2p-peer-address = " + strings.Join(otherPeers, "\np2p-peer-address = "),
		strings.Join(otherPeers, ","),
		"producer-name = " + strings.Join(names, "\nproducer-name = "),
		strings.Join(names, ","),
	}, nil)
}

func (b *BIOS) DispatchBootPublishHandoff() error {
	return b.dispatch("boot_publish_handoff", []string{
		b.EphemeralPrivateKey.PublicKey().String(),
		b.EphemeralPrivateKey.String(),
	}, nil)
}

func (b *BIOS) DispatchDone(operation string) error {
	return b.dispatch("done", []string{
		operation, // "join", "orchestrate", "boot"
	}, nil)
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
