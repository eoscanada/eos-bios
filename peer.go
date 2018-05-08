package bios

import (
	"fmt"

	"github.com/eoscanada/eos-bios/disco"
)

type Peer struct {
	DiscoveryFile IPFSRef
	DiscoveryLink IPNSRef // for ref
	Discovery     *disco.Discovery

	TotalWeight float64

	ClonedAccountName string
}

func (p *Peer) AccountName() string {
	if len(p.ClonedAccountName) != 0 {
		return p.ClonedAccountName
	}
	return string(p.Discovery.TargetAccountName)
}

func (p *Peer) String() string {
	return fmt.Sprintf("Account: % 15s   Org: % 30s   Weight: % 2.2f", p.AccountName(), p.Discovery.TargetAccountName, p.TotalWeight)
}

func (p *Peer) Columns() string {
	return fmt.Sprintf("%s | %s | %s | %.2f | %s", p.DiscoveryLink, p.AccountName(), p.Discovery.TargetAccountName, p.TotalWeight, p.DiscoveryFile)
}
