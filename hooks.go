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
	HookDef{"config_ready", "Dispatched when we are BIOS Node, and our keys and node config is ready. Should trigger a config update and a restart."},
	HookDef{"publish_kickstart_encrypted", "Dispatched with the contents of the encrypted Kickstart data, to be published to your social / web properties."},
	HookDef{"connect_to_bios", "Dispatched by ABPs with the decrypted contents of the Kickstart data.  Use this to initiate a connect from your BP node to the BIOS Node's p2p address."},
}

type HookDef struct {
	Key  string
	Desc string
}

func (b *BIOS) DispatchInit(genesisJSON string) error {
	return b.dispatch("init", []string{
		"genesis_json", genesisJSON,
	})
}

func (b *BIOS) DispatchConfigReady(genesisJSON, nodeName, publicKey, privateKey string, startProducing bool) error {
	return b.dispatch("config_ready", []string{
		"genesis_json", genesisJSON,
		"public_key", publicKey,
		"private_key", privateKey,
		"should_start_producing", fmt.Sprintf("%v", startProducing),
		"node_name", nodeName,
	})
}

func (b *BIOS) DispatchPublishKickstartEncrypted(kickstartData []byte) error {
	return b.dispatch("publish_kickstart_encrypted", []string{
		"data", string(kickstartData),
	})
}

func (b *BIOS) DispatchConnectToBIOS(p2pAddress, privateKeyUsed string) error {
	return b.dispatch("connect_to_bios", []string{
		"p2p_address", p2pAddress,
		"private_key_used", privateKeyUsed,
	})
}

// dispatch to both exec calls, and remote web hooks.
func (b *BIOS) dispatch(hookName string, data []string) error {
	conf := b.Config.Hooks[hookName]
	if conf == nil {
		return nil
	}

	fmt.Printf("Dispatching hook %q\n", hookName)

	if len(data)%2 != 0 {
		return fmt.Errorf("data should be pairs of key and values, cannot have %d elements", len(data))
	}

	if err := b.execCall(conf, data); err != nil {
		return err
	}

	if err := b.webhookCall(conf, data); err != nil {
		return err
	}

	if conf.Wait {
		fmt.Printf("Press ENTER to continue... ")
		_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	}

	return nil
}

func (b *BIOS) execCall(conf *HookConfig, data []string) error {
	if conf.Exec == "" {
		return nil
	}

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
	if conf.URL == "" {
		return nil
	}

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
