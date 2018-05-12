package main

import (
	"encoding/gob"
	"hash"
)

// HashGraph struct
type HashGraph struct {
	graph map[string]*Node
	// hash  hash.Hash64
}

// Node struct
type Node struct {
	sValue string
	iValue int
	uValue uint
	hash   hash.Hash64
	vert   [][]*Node
}

// NewHashGraph returns a pointer to a HashGraph object.
func NewHashGraph() *HashGraph {
	// Create the graphMap that will hold the hash graph.
	graphMap := make(map[string]*Node)
	// Encode the map into a byte stream.
	byteMap := gob.NewEncoder(nil)
	err := byteMap.Encode(graphMap)
	if err != nil {
		panic(err)
	}
	// Hash the graphMap to create a unique ID for the graph.
	return &HashGraph{
		graph: graphMap,
	}
}

// InsertNode inserts a node based on the value.
func (h *HashGraph) InsertNode(node Node) {
	// Check to see if the value already exists in the map.
	if h.graph[node.sValue] == nil {

	}
}

func main() {

}
