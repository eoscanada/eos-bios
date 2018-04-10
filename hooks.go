package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	shellwords "github.com/mattn/go-shellwords"
)

var configuredHooks = []HookDef{
	HookDef{"init", "Dispatch when we start the program."},
	HookDef{"start_bios_boot", "Dispatched when we are BIOS Node, and our keys and node config is ready. Should trigger a config update and a restart."},
	HookDef{"publish_kickstart_data", "Dispatched with the contents of the (usually encrypted) Kickstart data, to be published to your social / web properties."},
	HookDef{"connect_as_abp", "Dispatched by ABPs with the decrypted contents of the Kickstart data.  Use this to initiate a connect from your BP node to the BIOS Node's p2p address."},
	HookDef{"connect_as_participant", "Dispatched by all remaining participants (not BIOS Boot nor ABP) with the decrypted contents of the Kickstart data.  Use this to initiate a connect from your BP node to any of the Appointed Block Producers once they validated everything."},
	HookDef{"done", "When your process it done"},
}

type HookDef struct {
	Key  string
	Desc string
}

func (b *BIOS) DispatchInit() error {
	return b.dispatch("init", []string{}, nil)
}

func (b *BIOS) DispatchStartBIOSBoot(genesisJSON, publicKey, privateKey string) error {
	return b.dispatch("start_bios_boot", []string{
		"genesis_json", genesisJSON,
		"public_key", publicKey,
		"private_key", privateKey,
	}, nil)
}

func (b *BIOS) DispatchConnectAsABP(kickstart KickstartData, builtin func() error) error {
	return b.dispatch("connect_as_abp", []string{
		"p2p_address", kickstart.BIOSP2PAddress,
		"public_key_used", kickstart.PublicKeyUsed,
		"private_key_used", kickstart.PrivateKeyUsed,
		"genesis_json", kickstart.GenesisJSON,
	}, builtin)
}

func (b *BIOS) DispatchConnectAsParticipant(kickstart KickstartData, builtin func() error) error {
	return b.dispatch("connect_as_participant", []string{
		"p2p_address", kickstart.BIOSP2PAddress,
		"public_key_used", kickstart.PublicKeyUsed,
		"private_key_used", kickstart.PrivateKeyUsed,
		"genesis_json", kickstart.GenesisJSON,
	}, builtin)
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
	conf := b.Config.Hooks[hookName]
	if conf == nil {
		return nil
	}

	fmt.Printf("Dispatching hook %q\n", hookName)

	if len(data)%2 != 0 {
		return fmt.Errorf("data should be pairs of key and values, cannot have %d elements", len(data))
	}

	if conf.Exec != "" {
		if err := b.execCall(conf, data); err != nil {
			return err
		}
	}
	if conf.URL != "" {
		if err := b.webhookCall(conf, data); err != nil {
			return err
		}
	}
	if conf.Builtin {
		if f != nil {
			if err := f(); err != nil {
				return err
			}
		}
	}
	if conf.Wait {
		fmt.Printf("Press ENTER to continue... ")
		_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	}

	return nil
}

func (b *BIOS) execCall(conf *HookConfig, data []string) error {
	p := shellwords.NewParser()
	p.ParseEnv = true
	args, err := p.Parse(conf.Exec)
	if err != nil {
		return err
	}

	for i := 0; i < len(data); i += 2 {
		v := data[i+1]
		args = append(args, v)
	}

	var cmd *exec.Cmd
	if len(args) > 1 {
		cmd = exec.Command(args[0], args[1:]...)
	} else {
		cmd = exec.Command(args[0])
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()

	fmt.Printf("  Executing hook: %q\n", cmd.Args)

	return cmd.Run()
}

func (b *BIOS) webhookCall(conf *HookConfig, data []string) error {
	jsonBody, err := enc(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", conf.URL, jsonBody)
	if err != nil {
		return fmt.Errorf("NewRequest: %s", err)
	}

	// // Useful when debugging API calls
	// requestDump, err := httputil.DumpRequest(req, true)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(requestDump))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Do: %s", err)
	}
	defer resp.Body.Close()

	var cnt bytes.Buffer
	_, err = io.Copy(&cnt, resp.Body)
	if err != nil {
		return fmt.Errorf("Copy: %s", err)
	}

	if resp.StatusCode > 299 {
		return fmt.Errorf("status code=%d, body=%s", resp.StatusCode, cnt.String())
	}

	// fmt.Println("SERVER RESPONSE", cnt.String())

	return nil
}

func enc(v interface{}) (io.Reader, error) {
	if v == nil {
		return nil, nil
	}

	cnt, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	//fmt.Println("BODY", string(cnt))

	return bytes.NewReader(cnt), nil
}
