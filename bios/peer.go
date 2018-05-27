package bios

import (
	"fmt"
	"time"

	"github.com/eoscanada/eos-bios/bios/disco"
	eos "github.com/eoscanada/eos-go"
	"gonum.org/v1/gonum/graph"
)

type Peer struct {
	Discovery   *disco.Discovery
	UpdatedAt   time.Time
	TotalWeight int

	// ClonedAccountName string // A variation on the `Discovery`'s
}

// AccountName returns the variated account name (when cloned)
func (p *Peer) AccountName() string {
	// if len(p.ClonedAccountName) != 0 {
	// 	return p.ClonedAccountName
	// }
	// return string(p.Discovery.TargetAccountName)

	// TODO: implication of taking SeedNetwork rather than Target AccountName ?
	return string(p.Discovery.SeedNetworkAccountName)
}

// ID serves as a `graph.Node` implementation.
func (p *Peer) ID() int64 {
	id, _ := eos.StringToName(string(p.Discovery.SeedNetworkAccountName))
	return int64(id)
}

// Active flags peers if they published in the last 30
// minutes.. otherwise they're ignored, as to not stop the boot.
func (p *Peer) Active() bool {
	return p.UpdatedAt.After(time.Now().Add(-30 * time.Minute))
}

func (p *Peer) String() string {
	if p == nil {
		return "Account:nil"
	}
	if p.Discovery == nil {
		return "Discovery:nil"
	}

	return fmt.Sprintf("account=%s weight=%d", p.Discovery.SeedNetworkAccountName, p.TotalWeight)
}

func (p *Peer) Columns() string {
	active := ""
	if p.Active() {
		active = "A "
	}
	return fmt.Sprintf("%s | %s | %s%d | %d | %d", p.Discovery.SeedNetworkAccountName, p.Discovery.TargetAccountName, active, p.TotalWeight, p.Discovery.GMTOffset, p.Discovery.SeedNetworkLaunchBlock)
}

// PeerEdge is an internal structure that links two Discovery peers.
type PeerEdge struct {
	FromPeer *Peer
	ToPeer   *Peer
	PeerLink *disco.PeerLink
}

func (e *PeerEdge) From() graph.Node {
	return e.FromPeer
}

func (e *PeerEdge) To() graph.Node {
	return e.ToPeer
}

func (e *PeerEdge) Weight() float64 {
	return float64(e.PeerLink.Weight)
}
