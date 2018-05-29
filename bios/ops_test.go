package bios

import (
	"fmt"
	"testing"

	eos "github.com/eoscanada/eos-go"
	"github.com/stretchr/testify/assert"
)

func TestSnapshotDelegationAmounts(t *testing.T) {
	tests := []struct {
		balance  eos.Asset
		cpuStake eos.Asset
		netStake eos.Asset
	}{
		{
			eos.NewEOSAsset(10000), // 1.0 EOS
			eos.NewEOSAsset(2500),
			eos.NewEOSAsset(2500),
		},
		{
			eos.NewEOSAsset(100000), // 10.0 EOS
			eos.NewEOSAsset(2500),   // 0.25 EOS
			eos.NewEOSAsset(2500),   // 0.25 EOS
		},
		{
			eos.NewEOSAsset(105000), // 10.5 EOS
			eos.NewEOSAsset(2500),   // 0.25 EOS
			eos.NewEOSAsset(2500),   // 0.25 EOS
		},
		{
			eos.NewEOSAsset(107000), // 10.7 EOS
			eos.NewEOSAsset(3500),   // 0.35 EOS
			eos.NewEOSAsset(3500),   // 0.35 EOS
		},
		{
			eos.NewEOSAsset(120000), // 12.0 EOS
			eos.NewEOSAsset(10000),  // 0.25 + 0.75 EOS
			eos.NewEOSAsset(10000),  // 0.25 + 0.75 EOS
		},
		{
			eos.NewEOSAsset(99990000), // 9999.0 EOS
			eos.NewEOSAsset(49945000), // 4994.5 EOS
			eos.NewEOSAsset(49945000), // 4994.5 EOS, 10.0 EOS remaining :) yessir!
		},
	}

	for idx, test := range tests {
		cpuStake, netStake := splitSnapshotStakes(test.balance)
		assert.Equal(t, test.cpuStake, cpuStake, fmt.Sprintf("idx=%d", idx))
		assert.Equal(t, test.netStake, netStake, fmt.Sprintf("idx=%d", idx))
	}
}
