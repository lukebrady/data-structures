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
	g.Edges[u] = append(g.Edges[u], v)
	g.Edges[v] = append(g.Edges[v], u)
}

// PrintGraph prints out all of the nodes in the graph.
func (g *Graph) PrintGraph() {
	fmt.Println("Printing all edges in graph.")
	for u, adjacent := range g.Edges { // Nodes are labelled 0 to N-1.
		for _, v := range adjacent {
			// Edge exists from u to v.
			fmt.Printf("Edge: %d -> %d\n", u, v)
		}
	}
}

func main() {
	fmt.Println("Graphs!")
	graph := NewGraph(5)
	graph.AddNode(0, 1)
	graph.AddNode(1, 2)
	graph.AddNode(2, 0)
	graph.AddNode(1, 3)
	graph.AddNode(3, 4)
	graph.PrintGraph()
}
