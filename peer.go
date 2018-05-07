package bios

import (
	"fmt"
)

type Peer struct {
	DiscoveryFile IPFSRef
	DiscoveryLink IPNSRef // for ref
	Discovery     *Discovery

	TotalWeight float64

	ClonedAccountName string
}

func (p *Peer) AccountName() string {
	if len(p.ClonedAccountName) != 0 {
		return p.ClonedAccountName
	}
	return p.Discovery.EOSIOAccountName
}

func (p *Peer) String() string {
	if p == nil {
		return "Account:nil"
	}
	if p.Discovery == nil {
		return "Discovery:nil"
	}

	return fmt.Sprintf("Account: % 15s   Org: % 30s   Weight: % 2.2f", p.AccountName(), p.Discovery.OrganizationName, p.TotalWeight)
}
