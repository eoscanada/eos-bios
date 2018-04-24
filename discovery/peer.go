package discovery

import (
	"fmt"
)

type Peer struct {
	DiscoveryURL string
	Discovery    *Discovery

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
	return fmt.Sprintf("Account: % 15s   Org: % 30s   Weight: % 2.2f", p.AccountName(), p.Discovery.OrganizationName, p.TotalWeight)
}
