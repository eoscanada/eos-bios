package bios

import (
	"fmt"
	"testing"

	"github.com/eoscanada/eos-bios/bios/disco"
	eos "github.com/eoscanada/eos-go"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

func TestNetGraph(t *testing.T) {
	g := simple.NewDirectedGraph()
	d1 := newDisco(g, "a")
	d2 := newDisco(g, "b")
	d3 := newDisco(g, "c")
	d4 := newDisco(g, "d")
	d5 := newDisco(g, "e")
	d6 := newDisco(g, "f")
	d7 := newDisco(g, "g")
	d8 := newDisco(g, "h")

	addEdge(g, d1, d2)
	addEdge(g, d2, d3)
	addEdge(g, d3, d4)
	addEdge(g, d4, d1) // first cycle

	addEdge(g, d5, d6)
	addEdge(g, d6, d7)
	addEdge(g, d7, d8)
	addEdge(g, d8, d5) // second cycle

	addEdge(g, d3, d6) // link the two groups
	//	addEdge(g, d5, d4) // link the two groups both ways

	fmt.Printf("CYCLES:\n")
	for _, cycle := range topo.TarjanSCC(g) {
		fmt.Printf("Cycle: ")
		for _, el := range cycle {
			n := eos.NameToString(uint64(el.ID()))
			fmt.Printf("%s ", n)
		}
		fmt.Printf("\n")
	}

	fmt.Println("Nodes order:")
	for _, node := range g.Nodes() {
		fmt.Println("-", node.ID())
	}
}

func newDisco(g *simple.DirectedGraph, name string) *Peer {
	d := &Peer{Discovery: &disco.Discovery{SeedNetworkAccountName: eos.AN(name)}}
	g.AddNode(d)
	return d
}

func addEdge(g *simple.DirectedGraph, from, to graph.Node) {
	edge := g.NewEdge(from, to)
	g.SetEdge(edge)
}
