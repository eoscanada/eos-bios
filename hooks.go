package main

import (
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

type HookInit struct{}

type HookConfigReady struct {
	GenesisJSON string `json:"genesis_json"`
	PublicKey   string `json:"public_key"`
	PrivateKey  string `json:"private_key"`
}

type HookPublishKickstartEncrypted struct {
	Data []byte
}

type HookConnectToBIOS struct {
	P2PAddress     string `json:"p2p_address"`
	PrivateKeyUsed string `json:"private_key_used"`
}

type HookPublishKickstartPublic struct {
	P2PAddress     string `json:"p2p_address"`
	PrivateKeyUsed string `json:"private_key_used"`
}

func (b *BIOS) DispatchInit() error {
	return b.dispatch(b.Config.Hooks["init"], &HookInit{})
}

func (b *BIOS) DispatchConfigReady(genesisJSON string, publicKey string, privateKey string) error {
	return b.dispatch(b.Config.Hooks["config_ready"], &HookConfigReady{
		GenesisJSON: genesisJSON,
		PublicKey:   publicKey,
		PrivateKey:  privateKey,
	})
}

func (b *BIOS) DispatchPublishKickstartEncrypted(kickstartData []byte) error {
	return b.dispatch(b.Config.Hooks["publish_kickstart_encrypted"], &HookPublishKickstartEncrypted{
		Data: kickstartData,
	})
}

func (b *BIOS) DispatchConnectToBIOS(p2pAddress, privateKeyUsed string) error {
	return b.dispatch(b.Config.Hooks["connect_to_bios"], &HookConnectToBIOS{
		P2PAddress:     p2pAddress,
		PrivateKeyUsed: privateKeyUsed,
	})
}

// dispatch to both exec calls, and remote web hooks.
func (b *BIOS) dispatch(conf *HookConfig, data interface{}) error {
	if conf == nil {
		return nil
	}

	if err := b.execCall(conf, data); err != nil {
		return err
	}

	if err := b.webhookCall(conf, data); err != nil {
		return err
	}

	return nil
}

func (b *BIOS) execCall(conf *HookConfig, data interface{}) error {
	if conf.execTemplate == nil {
		return nil
	}

	var buf bytes.Buffer
	if err := conf.execTemplate.Execute(&buf, data); err != nil {
		return err
	}

	p := shellwords.NewParser()
	p.ParseEnv = true
	args, err := p.Parse(buf.String())
	if err != nil {
		return err
	}

	var cmd *exec.Cmd
	if len(args) > 1 {
		cmd = exec.Command(args[0], args[1:]...)
	} else {
		cmd = exec.Command(args[0])
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()
	return cmd.Start()
}

func (b *BIOS) webhookCall(conf *HookConfig, data interface{}) error {
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
