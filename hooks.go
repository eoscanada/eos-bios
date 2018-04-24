package bios

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/eoscanada/eos-bios/discovery"
	shellwords "github.com/mattn/go-shellwords"
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
	return b.dispatch("start_bios_boot", []string{
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

	privKey := ""
	if b.Config.Peer.BlockSigningPrivateKey != nil {
		privKey = b.Config.Peer.BlockSigningPrivateKey.String()
	}

	return b.dispatch("connect_as_abp", []string{
		"p2p_address", kickstart.BIOSP2PAddress,
		"public_key", peerDefs[0].Discovery.EOSIOABPSigningKey.String(),
		"private_key", privKey,
		"genesis_json", kickstart.GenesisJSON,
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
