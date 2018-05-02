package bios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitcoinClock(t *testing.T) {
	hash, err := bitcoinPollMethodBlockchainInfo(520895)
	assert.NoError(t, err)
	assert.Equal(t, "000000000000000000275203ae47eb5ad81b13fe40e7456927894b94f295f7fb", hash)

	hash, err = bitcoinPollMethodBlockchainInfo(10000000)
	assert.NoError(t, err)
	assert.Equal(t, "", hash)

	hash, err = bitcoinPollMethodBlockExplorer(520895)
	assert.NoError(t, err)
	assert.Equal(t, "000000000000000000275203ae47eb5ad81b13fe40e7456927894b94f295f7fb", hash)

	hash, err = bitcoinPollMethodBlockExplorer(10000000)
	assert.NoError(t, err)
	assert.Equal(t, "", hash)
}
