package bios

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetMeshList(t *testing.T) {
	ints := makeRange(0, 35)

	head := "|NODE|"
	out := ""

	for idx := range ints {
		list := GetMeshList(len(ints), idx)

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
func TestGetMeshListToJson(t *testing.T) {
	ints := makeRange(0, 35)

	allNodes := []*node{}

	for idx := range ints {
		list := GetMeshList(len(ints), idx)
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
	fmt.Println(string(res))

}

type node struct {
	Name  string   `json:"name"`
	Peers []string `json:"peers"`
}
