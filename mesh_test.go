package bios

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

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
// `go test -timeout 30s github.com/eoscanada/eos-bios -run ^TestGetMeshListToJson`
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
