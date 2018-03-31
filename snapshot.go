package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/eosioca/eosapi"
	"github.com/eosioca/eosapi/ecc"
)

type Snapshot []SnapshotLine

type SnapshotLine struct {
	EthereumAddress string
	EOSPublicKey    ecc.PublicKey
	Balance         eos.Asset
}

func NewSnapshot(filename string) (out Snapshot, err error) {
	fl, err := os.Open(filename)
	if err != nil {
		return
	}

	reader := csv.NewReader(fl)
	allRecords, err := reader.ReadAll()
	if err != nil {
		return
	}

	//fmt.Println("ALL records", allRecords)

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
