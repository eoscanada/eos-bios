package bios

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/eoscanada/eos-bios/bios/disco"
	eos "github.com/eoscanada/eos-go"
	"github.com/stretchr/testify/assert"
)

func TestShuffling(t *testing.T) {
	tests := []struct {
		numPeers int
		seed     int64
		out      string
	}{
		{10, 0, "p2,p1,p0"}, // from p0 to p9
		{10, 1, "p1,p2,p0"},
		{10, 777, "p0,p2,p1"},
		{20, 567, "p2,p0,p1,p3,p4"},
	}

	for _, test := range tests {
		var peers []*Peer
		for i := 0; i < test.numPeers; i++ {
			peers = append(peers, &Peer{Discovery: &disco.Discovery{SeedNetworkAccountName: eos.AccountName(fmt.Sprintf("p%d", i))}})
		}
		b := &BIOS{
			ShuffledProducers: peers,
			RandSource:        rand.NewSource(test.seed),
		}

		b.shuffleProducers()

		expectedPeers := strings.Split(test.out, ",")
		for idx, el := range expectedPeers {
			assert.Equal(t, el, b.ShuffledProducers[idx].AccountName(), fmt.Sprintf("Seed %d", test.seed))
		}
	}
}
