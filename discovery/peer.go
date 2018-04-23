package discovery

type Peer struct {
	DiscoveryURL string
	TotalWeight  float64
	Edges        []string
	Discovery    *Discovery
}
