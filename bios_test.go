package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"
)

func TestIsABP(t *testing.T) {
	bios := testBIOS(t, `
producers:
- eosio_account_name: mama
- eosio_account_name: papa
- eosio_account_name: teen
`, `
producer:
  my_account: mama
`)
	assert.True(t, bios.AmIBootNode())
	assert.False(t, bios.IsAppointedBlockProducer("mama"))
	assert.True(t, bios.IsAppointedBlockProducer("papa"))
	assert.True(t, bios.IsAppointedBlockProducer("teen"))
}

func testBIOS(t *testing.T, launchyaml string, config string) *BIOS {
	b := &BIOS{}
	err := yaml.Unmarshal([]byte(config), &b.Config)
	require.NoError(t, err)

	err = yaml.Unmarshal([]byte(launchyaml), &b.LaunchData)
	require.NoError(t, err)

	b.Config.NoShuffle = true

	require.NoError(t, b.ShuffleProducers([]byte{}))

	return b
}
