package main

import "fmt"

// Node contains data that will be used in the weighted graph.
type Node struct {
	value   int
	weight  int
	message string
}

// WeightedGraph struct.
type WeightedGraph struct {
	numNodes int
	edges    [][]Node
}

// NewWeightedGraph creates a weighted graph object.
func NewWeightedGraph(nodes int) *WeightedGraph {
	return &WeightedGraph{
		numNodes: nodes,
		edges:    make([][]Node, nodes),
	}
}

// AddNode adds a weighted node to the graph.
func (g *WeightedGraph) AddNode(node1, node2 Node) {
	// Append the weighted node to the graph.
	g.edges[node1.value] = append(g.edges[node1.value], node2)
	//
	g.edges[node2.value] = append(g.edges[node2.value], node1)
}

// PrintGraph prints out all of the nodes in the graph.
func (g *WeightedGraph) PrintGraph() {
	fmt.Println("Printing all edges in graph.")
	for u, adjacent := range g.edges { // Nodes are labelled 0 to N-1.
		for _, node := range adjacent {
			// Edge exists from u to v.
			fmt.Printf("Edge: %d -> %d\n", u, node.value)
		}
	}
}

// PrintGraphMessages prints out all of the nodes in the graph.
func (g *WeightedGraph) PrintGraphMessages() {
	fmt.Println("Printing all edges in graph.")
	for _, adjacent := range g.edges { // Nodes are labelled 0 to N-1.
		for _, node := range adjacent {
			// Edge exists from u to v.
			fmt.Printf("Edge: %d -> %s\n", node.value, node.message)
		}
	}
}

// NewNode creates a new weighted node object.
func NewNode(value, weight int, message string) *Node {
	return &Node{
		value:   value,
		weight:  weight,
		message: message,
	}
}

func main() {
	node1 := NewNode(2, 12, "Hello there I am node1.")
	node2 := NewNode(1, 3, "Hello there I am node2.")
	wGraph := NewWeightedGraph(10)
	wGraph.AddNode(*node1, *node2)
	wGraph.PrintGraph()
	wGraph.PrintGraphMessages()
}
