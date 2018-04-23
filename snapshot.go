package bios

import (
	"bytes"
	"encoding/csv"
	"fmt"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
)

type Snapshot []SnapshotLine

type SnapshotLine struct {
	EthereumAddress string
	EOSPublicKey    ecc.PublicKey
	Balance         eos.Asset
}

func NewSnapshot(content []byte) (out Snapshot, err error) {
	reader := csv.NewReader(bytes.NewBuffer(content))
	allRecords, err := reader.ReadAll()
	if err != nil {
		return
	}

	for _, el := range allRecords {
		if len(el) != 3 {
			return nil, fmt.Errorf("should have 3 elements per line")
		}

		newAsset, err := eos.NewEOSAssetFromString(el[2])
		if err != nil {
			return out, err
		}

		pubKey, err := ecc.NewPublicKey(el[1])
		if err != nil {
			return out, err
		}

		out = append(out, SnapshotLine{el[0], pubKey, newAsset})
	}

	return
}
