package bios

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"testing"

	"github.com/eoscanada/eos-bios/bios/disco"
	eos "github.com/eoscanada/eos-go"
	"github.com/stretchr/testify/assert"
)

type node struct {
	Name  string   `json:"name"`
	Peers []string `json:"peers"`
}

func makeRange(count int) []int {
	a := make([]int, count)
	for i := range a {
		a[i] = i
	}
	return a
}
func TestGetMeshList(t *testing.T) {
	ints := makeRange(35)

	head := "|NODE|"
	out := ""

	for idx := range ints {
		list := getPeerIndexesToMeshWith(len(ints), idx)

		head += fmt.Sprintf("%03d|", idx)
		res := ""
		for _, nodeNum := range ints {
			if list[nodeNum] {
				res += " x |"
			} else {
				res += "   |"
			}
		}
		out += fmt.Sprintf("|%04d|%s\n", idx, res)
	}
	fmt.Println(head)
	fmt.Println(out)

}

// To visually test the meshing, run
// `go test -timeout 30s github.com/eoscanada/eos-bios/bios -run ^TestGetMeshListToJson`
// `cd test-data/mesh/`
// `python3 -m http.server`
// open `http://localhost:8000?count=` + a number of nodes you want to visualize
func TestGetMeshListToJson(t *testing.T) {
	tests := []int{1, 2, 3, 4, 7, 12, 16, 21, 35, 56, 85, 121}

	for _, numNodes := range tests {
		ints := makeRange(numNodes)

		allNodes := []*node{}

		for idx := range ints {
			list := getPeerIndexesToMeshWith(len(ints), idx)
			peers := []string{}
			for index := range list {
				peers = append(peers, fmt.Sprintf("%d", index))
			}
			node := &node{
				Name:  fmt.Sprintf("%d", idx),
				Peers: peers,
			}
			allNodes = append(allNodes, node)

		}
		res, _ := json.Marshal(allNodes)
		err := ioutil.WriteFile(fmt.Sprintf("./test-data/mesh/flare_%d.json", numNodes), []byte(res), 0644)
		assert.NoError(t, err)

	}

}

func TestGetPeersForBootNode(t *testing.T) {
	tests := []struct {
		numPeers int
		seed     int64
		out      string
	}{
		{1, 1, "p0"},                                                                                             // 1
		{3, 1, "p0,p1,p2"},                                                                                       // 3
		{10, 1, "p0,p1,p2,p3,p4,p5,p6,p7,p8,p9"},                                                                 // 10
		{40, 1, "p10,p23,p3,p39,p6,p7,p20,p12,p0,p4,p33,p8,p14,p26,p18,p37,p17,p11,p24,p9,p30,p5,p27,p21,p31"},   // 25
		{40, 2, "p31,p3,p23,p5,p16,p0,p15,p35,p38,p21,p14,p19,p32,p17,p20,p11,p7,p6,p37,p25,p34,p18,p9,p30,p26"}, // 25
		{60, 1, "p10,p12,p9,p3,p1,p7,p6,p15,p16,p13,p8,p5,p18,p4,p19,p43,p28,p24,p32,p23,p53,p54,p58,p46,p47"},   // 25
	}

	for _, test := range tests {
		var peers []*Peer
		for i := 0; i < test.numPeers; i++ {
			peers = append(peers, &Peer{Discovery: &disco.Discovery{SeedNetworkAccountName: eos.AccountName(fmt.Sprintf("p%d", i))}})
		}
		b := &BIOS{}

		listOfPeers := b.getPeersForBootNode(peers, rand.NewSource(test.seed))
		expectedPeers := strings.Split(test.out, ",")
		for idx, el := range expectedPeers {
			assert.Equal(t, el, listOfPeers[idx].AccountName())
		}

	}
}
