package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsABP(t *testing.T) {
	bios := testBIOS(t, `
producers:
- account_name: mama
- account_name: papa
- account_name: teen
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
	err := yamlUnmarshal([]byte(config), &b.Config)
	require.NoError(t, err)

	err = yamlUnmarshal([]byte(launchyaml), &b.LaunchData)
	require.NoError(t, err)

	b.Config.NoShuffle = true

	require.NoError(t, b.ShuffleProducers([]byte{}, time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC)))

	return b
}
