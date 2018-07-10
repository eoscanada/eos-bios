package bios

import (
	"fmt"
	"os"
	"os/exec"
)

func (b *BIOS) DispatchBootNode(genesisJSON, publicKey, privateKey string) error {
	return b.dispatch("boot", []string{
		genesisJSON,
		publicKey,
		privateKey,
	}, nil)
}

// dispatch to both exec calls, and remote web hooks.
func (b *BIOS) dispatch(hookName string, args []string, f func() error) error {
	b.Log.Printf("---- BEGIN HOOK %q ----\n", hookName)

	executable := fmt.Sprintf("./%s.sh", hookName)

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

	b.Log.Printf("---- END HOOK %q ----\n", hookName)

	return nil
}
