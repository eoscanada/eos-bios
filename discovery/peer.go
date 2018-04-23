package discovery

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
