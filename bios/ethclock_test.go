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

func TestEthereumClock(t *testing.T) {
	hash, err := etherscanPollMethod(5544735)
	assert.NoError(t, err)
	assert.Equal(t, "0552171fbbd84d6fc7ee2371b2de61371d9a291aa5ce96521d4f595363f7eee9", hash)

	hash, err = etherscanPollMethod(1000000000)
	assert.NoError(t, err)
	assert.Equal(t, "", hash)
}
