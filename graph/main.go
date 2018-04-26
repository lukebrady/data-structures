package main

import (
	"fmt"
)

// Graph is the base type.
type Graph struct {
	NumNodes int
	Edges    [][]int
}

// NewGraph returns a pointer to a graph object.
func NewGraph(nodes int) *Graph {
	return &Graph{
		NumNodes: nodes,
		Edges:    make([][]int, nodes),
	}
}

// AddNode adds a node to a created Graph.
func (g *Graph) AddNode(u, v int) {

}

func main() {
	fmt.Println("Graphs!")
}
