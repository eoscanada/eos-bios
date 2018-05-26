package bios

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

// TODO: update with latest GenesisJSON with the basic parameters...
type GenesisJSON struct {
	InitialTimestamp string `json:"initial_timestamp"`
	InitialKey       string `json:"initial_key"`
}

func readGenesisData(text string, ipfs *IPFS) (out *GenesisJSON, err error) {
	// try base64 encoded genesis
	text = strings.TrimSpace(text)
	if strings.HasPrefix(text, "ey") {
		// base64 decode
		decoded, err := base64.RawStdEncoding.DecodeString(text)
		if err != nil {
			return nil, err
		}
		text = string(decoded)
	}

	if strings.Contains(text, "/ipfs/") {
		// fetch from IPFS and decode
		cnt, err := ipfs.Get(text)
		if err != nil {
			return nil, err
		}

		text = string(cnt)
	}

	if strings.HasPrefix(text, "{") {
		err = json.Unmarshal([]byte(text), &out)
		return
	}

	return nil, errors.New("invalid genesis data, not base64-encoded JSON, not JSON, not an ipfs link, what was that anyway?")
}
