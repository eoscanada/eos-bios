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

type HookInit struct{}

type HookConfigReady struct{}

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
	conf := b.Config.Hooks.Init
	return b.dispatch(conf, &HookInit{})
}

func (b *BIOS) DispatchConfigReady() error {
	conf := b.Config.Hooks.ConfigReady
	return b.dispatch(conf, &HookConfigReady{})
}

func (b *BIOS) DispatchPublishKickstartEncrypted(kickstartData []byte) error {
	conf := b.Config.Hooks.PublishKickstartEncrypted
	return b.dispatch(conf, &HookPublishKickstartEncrypted{
		Data: kickstartData,
	})
}

func (b *BIOS) DispatchConnectToBIOS(p2pAddress, privateKeyUsed string) error {
	conf := b.Config.Hooks.ConnectToBIOS
	return b.dispatch(conf, &HookConnectToBIOS{
		P2PAddress:     p2pAddress,
		PrivateKeyUsed: privateKeyUsed,
	})
}

func (b *BIOS) DispatchPublishKickstartPublic(p2pAddress, privateKeyUsed string) error {
	conf := b.Config.Hooks.PublishKickstartPublic
	return b.dispatch(conf, &HookPublishKickstartPublic{
		P2PAddress:     p2pAddress,
		PrivateKeyUsed: privateKeyUsed,
	})
}

// dispatch to both `exec` calls, and remote web hooks.
func (b *BIOS) dispatch(conf HookConfig, data interface{}) error {
	if err := b.execCall(conf, data); err != nil {
		return err
	}

	if err := b.webhookCall(conf, data); err != nil {
		return err
	}

	return nil
}

func (b *BIOS) execCall(conf HookConfig, data interface{}) error {
	execTpl, err := conf.parseTemplate()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := execTpl.Execute(&buf, data); err != nil {
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

func (b *BIOS) webhookCall(conf HookConfig, data interface{}) error {
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
