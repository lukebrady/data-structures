package main

// Graph datastructure
type Graph struct {
	size  uint
	edges map[string][]Node
}

// Node that holds data within the graph.
type Node struct {
	name   string
	weight uint
	data   []byte
}

// NewGraph produces a new graph object.
func NewGraph() *Graph {
	var edgeArr map[string][]Node
	return &Graph{
		size:  0,
		edges: edgeArr,
	}
}

// AddNode adds a data node the graph.
func (g *Graph) AddNode(name string, neighbors []Node, weight uint) {
	// Check to see if the node exists in the graph.
}

func main() {

}
