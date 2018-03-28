package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Snapshot []SnapshotLine

type SnapshotLine struct {
	EthereumAddress string
	EOSPublicKey    string
	Balance         string
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

	for _, el := range allRecords {
		if len(el) != 3 {
			return nil, fmt.Errorf("should have 3 elements per line")
		}

		out = append(out, SnapshotLine{el[0], el[1], el[2]})
	}

	return
}
