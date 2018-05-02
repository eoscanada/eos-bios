package bios

import (
	"fmt"
	"testing"
)

func TestGetMeshList(t *testing.T) {
	ints := makeRange(1, 36)

	head := "|NODE|"
	out := ""

	for idx := range ints {
		list := GetMeshList(len(ints), idx)

		head += fmt.Sprintf("%03d|", idx+1)
		res := ""
		for _, nodeNum := range ints {
			if list[nodeNum] {
				res += " x |"
			} else {
				res += "   |"
			}
		}
		out += fmt.Sprintf("|%04d|%s\n", idx+1, res)
	}
	fmt.Println(head)
	fmt.Println(out)

}
