package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WebhookInit struct{}

type WebhookConfigReady struct{}

type WebhookPublishKickstartEncrypted struct {
	Data []byte
}

type WebhookConnectToBIOS struct {
	P2PAddress     string `json:"p2p_address"`
	PrivateKeyUsed string `json:"private_key_used"`
}

type WebhookPublishKickstartPublic struct {
	P2PAddress     string `json:"p2p_address"`
	PrivateKeyUsed string `json:"private_key_used"`
}

func (b *BIOS) DispatchInit() error {
	conf := b.Config.Webhooks.Init
	return webhookCall(conf.URL, &WebhookInit{})
}

func (b *BIOS) DispatchConfigReady() error {
	conf := b.Config.Webhooks.ConfigReady
	return webhookCall(conf.URL, &WebhookConfigReady{})
}

func (b *BIOS) DispatchPublishKickstartEncrypted(kickstartData []byte) error {
	conf := b.Config.Webhooks.PublishKickstartEncrypted
	return webhookCall(conf.URL, &WebhookPublishKickstartEncrypted{
		Data: kickstartData,
	})
}

func (b *BIOS) DispatchConnectToBIOS(p2pAddress, privateKeyUsed string) error {
	conf := b.Config.Webhooks.ConnectToBIOS
	return webhookCall(conf.URL, &WebhookConnectToBIOS{
		P2PAddress:     p2pAddress,
		PrivateKeyUsed: privateKeyUsed,
	})
}

func (b *BIOS) DispatchPublishKickstartPublic(p2pAddress, privateKeyUsed string) error {
	conf := b.Config.Webhooks.PublishKickstartPublic
	return webhookCall(conf.URL, &WebhookPublishKickstartPublic{
		P2PAddress:     p2pAddress,
		PrivateKeyUsed: privateKeyUsed,
	})
}

func webhookCall(endpoint string, data interface{}) error {
	if endpoint == "" {
		return nil
	}

	jsonBody, err := enc(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", endpoint, jsonBody)
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
