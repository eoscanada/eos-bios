package bios

import (
	"math"
	"math/rand"
	"time"
)

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
	myPosition := -1

	meshableProducers := b.meshableShuffledProducers()

	for idx, peer := range meshableProducers {
		mySeedAccount := b.Network.MyPeer.Discovery.SeedNetworkAccountName
		if mySeedAccount == peer.Discovery.SeedNetworkAccountName {
			myPosition = idx
		}
	}
	if myPosition != -1 {
		peerIDs := getPeerIndexesToMeshWith(len(meshableProducers), myPosition)
		for idx, peer := range meshableProducers {
			p2pAddr := peer.Discovery.TargetP2PAddress
			if peerIDs[idx] && !otherPeersMap[p2pAddr] {
				otherPeers = append(otherPeers, p2pAddr)
				otherPeersMap[p2pAddr] = true
			}
		}
	}
	return otherPeers
}

func (b *BIOS) getPeersForBootNode(orderedPeers []*Peer, randSource rand.Source) (out []*Peer) {
	r := rand.New(randSource)

	if len(orderedPeers) < 26 {
		return orderedPeers
	}
	if len(orderedPeers) > 50 {
		top := shuffle(orderedPeers[:20], 15, r)
		part2 := shuffle(orderedPeers[20:45], 5, r)
		part3 := shuffle(orderedPeers[45:], 5, r)
		top = append(top, part2...)
		return append(top, part3...)

	}
	return shuffle(orderedPeers, 25, r)
}

func shuffle(slice []*Peer, count int, r rand.Source) []*Peer {
	ret := make([]*Peer, count)
	for i := 0; i < count; i++ {
		randIndex := r.Int63() % int64(len(slice))
		ret[i] = slice[randIndex]
		slice = append(slice[:randIndex], slice[randIndex+1:]...)
	}
	return ret
}

func (b *BIOS) someTopmostPeersAddresses() []string {
	meshableProducers := b.meshableShuffledProducers()

	listOfPeers := b.getPeersForBootNode(meshableProducers, rand.NewSource(time.Now().UTC().UnixNano()))
	otherPeers := []string{}
	for _, peer := range listOfPeers {
		otherPeers = append(otherPeers, peer.Discovery.TargetP2PAddress)
	}
	return otherPeers
}

func (b *BIOS) meshableShuffledProducers() []*Peer {
	var meshableProducers []*Peer
	for _, peer := range b.ShuffledProducers {
		if peer.Discovery.TargetP2PAddress == "none" {
			continue
		}

		if !peer.Active() {
			continue
		}

		meshableProducers = append(meshableProducers, peer)
	}
	return meshableProducers
}
