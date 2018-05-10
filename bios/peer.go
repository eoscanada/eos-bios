package bios

import (
	"fmt"

	"github.com/eoscanada/eos-bios/bios/disco"
)

type Peer struct {
	Discovery *disco.Discovery

	TotalWeight int

	ClonedAccountName string
}

func (p *Peer) AccountName() string {
	if len(p.ClonedAccountName) != 0 {
		return p.ClonedAccountName
	}
	return string(p.Discovery.TargetAccountName)
}

func (p *Peer) String() string {
	if p == nil {
		return "Account:nil"
	}
	if p.Discovery == nil {
		return "Discovery:nil"
	}

	return fmt.Sprintf("account=%s weight=%d", p.AccountName(), p.TotalWeight)
}

func (p *Peer) Columns() string {
	return fmt.Sprintf("%s | %s | %d | %d | %d", p.Discovery.SeedNetworkAccountName, p.Discovery.TargetAccountName, p.TotalWeight, p.Discovery.GMTOffset, p.Discovery.SeedNetworkLaunchBlock)
}
