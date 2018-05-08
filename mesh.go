package bios

import "math"

func getPeerIndexesToMeshWith(total, myPos int) map[int]bool {
	list := map[int]bool{}
	firstNeighbour := (myPos + 1) % total

	list[firstNeighbour] = true

	for i := 0; i < numConnectionsRequired(total); i++ {
		nextNeighbour := int(math.Pow(2.0, float64(i+2))) + myPos
		if nextNeighbour >= total {
			nextNeighbour = (nextNeighbour) % total
		}
		list[nextNeighbour] = true
	}
	return list
}

func numConnectionsRequired(numberOfNodes int) int {
	return int(math.Ceil(math.Sqrt(float64(numberOfNodes))))
}

func (b *BIOS) computeMyMeshP2PAddresses() []string {
	otherPeers := []string{}
	otherPeersMap := map[string]bool{}
	if len(b.MyPeers) > 0 {
		myFirstPeer := b.MyPeers[0]
		myPosition := -1
		for idx, peer := range b.ShuffledProducers {
			if myFirstPeer.AccountName() == peer.AccountName() {
				myPosition = idx
			}
		}
		if myPosition != -1 {
			peerIDs := getPeerIndexesToMeshWith(len(b.ShuffledProducers), myPosition)
			for idx, peer := range b.ShuffledProducers {
				p2pAddr := peer.Discovery.TargetP2PAddress
				if peerIDs[idx] && !otherPeersMap[p2pAddr] {
					otherPeers = append(otherPeers, p2pAddr)
					otherPeersMap[p2pAddr] = true
				}
			}
		}
	}
	return otherPeers
}

func (b *BIOS) allExceptBootP2PAddresses() (out []string) {
	for _, el := range b.Network.discoveredPeers {
		if el.Discovery.SeedNetworkAccountName == b.Network.MyPeer.Discovery.SeedNetworkAccountName {
			continue
		}

		// TODO: also skip boot node, as he wasn't participating in the launch

		out = append(out, el.Discovery.TargetP2PAddress)
	}
	return
}

func (b *BIOS) someTopmostPeersAddresses() []string {
	// TODO: refine this algo..
	// connect to some randomly, but more of the top-most
	otherPeers := []string{}
	for idx, peer := range b.ShuffledProducers {
		if idx > 5 {
			return otherPeers
		}
		otherPeers = append(otherPeers, peer.Discovery.TargetP2PAddress)
	}
	return otherPeers
}
