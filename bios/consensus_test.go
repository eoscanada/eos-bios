package bios

import (
	"testing"

	"github.com/eoscanada/eos-bios/bios/disco"
	"github.com/stretchr/testify/assert"
)

func TestConsensusContents(t *testing.T) {
	myDisco := &disco.Discovery{SeedNetworkAccountName: AN("p1")}
	net := NewNetwork("/tmp/consensus-test-cache", myDisco, nil, "eosio.disco", nil)
	net.allNodesFetchFunc = func() error {
		p1 := &Peer{
			Discovery: &disco.Discovery{
				SeedNetworkAccountName: AN("p1"),
				SeedNetworkPeers: []*disco.PeerLink{
					&disco.PeerLink{Account: AN("p2")},
					&disco.PeerLink{Account: AN("p3")},
				},
				TargetContents: []disco.ContentRef{
					disco.ContentRef{
						Name: "file1",
						Ref:  "/ipfs/file1",
					},
					disco.ContentRef{
						Name: "file2",
						Ref:  "/ipfs/file2-modified",
					},
				},
			},
		}
		net.allNodes.AddNode(p1)

		p2 := &Peer{
			Discovery: &disco.Discovery{
				SeedNetworkAccountName: AN("p2"),
				SeedNetworkPeers: []*disco.PeerLink{
					&disco.PeerLink{Account: AN("p1")},
					&disco.PeerLink{Account: AN("p3")},
				},
				TargetContents: []disco.ContentRef{
					disco.ContentRef{
						Name: "file1",
						Ref:  "/ipfs/file1-modified",
					},
					disco.ContentRef{
						Name: "file2",
						Ref:  "/ipfs/file2",
					},
					disco.ContentRef{
						Name: "file3",
						Ref:  "/ipfs/file3",
					},
				},
			},
		}
		net.allNodes.AddNode(p2)

		p3 := &Peer{
			Discovery: &disco.Discovery{
				SeedNetworkAccountName: AN("p3"),
				SeedNetworkPeers: []*disco.PeerLink{
					&disco.PeerLink{Account: AN("p1")},
					&disco.PeerLink{Account: AN("p2")},
				},
				TargetContents: []disco.ContentRef{
					disco.ContentRef{
						Name: "file1",
						Ref:  "/ipfs/file1",
					},
					disco.ContentRef{
						Name: "file2",
						Ref:  "/ipfs/file2",
					},
				},
			},
		}
		net.allNodes.AddNode(p3)

		return nil
	}
	assert.NoError(t, net.UpdateGraph())

	orderedPeers := net.OrderedPeers(net.MyNetwork())
	files := ComputeContentsAgreement(orderedPeers)

	assert.Len(t, files.FilesMap["file1"], 2)
	assert.Len(t, files.FilesMap["file2"], 2)
	assert.Len(t, files.FilesMap["file3"], 1)
	assert.Equal(t, 1, files.FilesMap["file1"]["/ipfs/file1"])
	assert.Equal(t, 2, files.FilesMap["file1"]["/ipfs/file1-modified"])

	assert.Equal(t, []string{"1", "2", "2"}, ComputePeerContentsColumn(files, orderedPeers))
}
