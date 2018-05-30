package bios

import (
	"testing"

	"github.com/eoscanada/eos-bios/bios/disco"
	"github.com/stretchr/testify/assert"
)

func TestBootOptInFilter(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"ABcDeFGhIJ", "ABDFGcehIJ"},
		{"ABCDEFGHIJ", "ABCDEFGHIJ"},
		{"abcdeFGHIJ", "FGHIJabcde"},
		{"abcdefGHIJ", "GHIJabcdef"},
		{"abcdefGHIJKLM", "GHIJKabcdefLM"},
		{"abcdefghIJ", "IJabcdefgh"},
		{"abcde", "abcde"},
		{"abC", "Cab"},
		{"abcdeF", "Fabcde"},
		{"abcdeFGHI", "FGHIabcde"},
		{"AbCdEfGhIjKlM", "ACEGIbdfhjKlM"},
	}

	for _, test := range tests {
		in := bootFilterPeers(test.in)
		out := bootFilterPeers(test.out)

		assert.Equal(t, out, bootOptInFilter(in))
	}
}

func bootFilterPeers(in string) []*Peer {
	var out []*Peer
	for _, c := range in {
		block := 0
		if c < 92 {
			block = 1
		}
		out = append(out, &Peer{Discovery: &disco.Discovery{
			SeedNetworkAccountName: AN(string(c)),
			SeedNetworkLaunchBlock: uint64(block),
		}})
	}
	return out
}
