package main

import (
	"os"

	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

// Define a custom node type that implements the graph.Node interface.
type Node int64

func (n Node) ID() int64 { return int64(n) }

func main() {
	g := simple.NewDirectedGraph()

	n1 := Node(1)
	n2 := Node(2)
	g.AddNode(n1)
	g.AddNode(n2)
	g.SetEdge(g.NewEdge(n1, n2))

	data, err := dot.Marshal(g, "MyGraph", "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("output.dot", data, 0o644)
}
