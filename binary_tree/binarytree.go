package main

import (
	"fmt"
)

// Node is used to define a node within a tree.
type Node struct {
	left  *Node
	right *Node
	data  int
}

// Tree is the tree type that is returned by NewTree.
type Tree struct {
	depth  int
	secret string
	root   *Node
}

// NewTree creates a new tree and returns it.
func NewTree(secrets string) *Tree {
	return &Tree{
		depth:  0,
		secret: secrets,
		root:   nil,
	}
}

// AddNode adds an int node to the root.
func (t *Tree) AddNode(data int) {
	node := &Node{data: data}
	if t.root == nil {
		t.root = node
		fmt.Printf("Root node is %d.\n", t.root.data)
		t.depth++
	} else {
		place := t.root
		for place.left != nil ||
			place.right != nil {
			if place.data > node.data {
				place = place.left
			} else if place.data < node.data {
				place = place.right
			}
		}
		fmt.Printf("Placing %d into the tree.\n", node.data)
		place = node
		t.depth++
	}
}

// GetDepth returns the depth of the tree.
func (t *Tree) GetDepth() int {
	return t.depth
}

func main() {
	tree := NewTree("Test")
	tree.AddNode(5)
	tree.AddNode(6)
	tree.AddNode(2)
	fmt.Printf("Tree depth is %d.\n", tree.GetDepth())
}
