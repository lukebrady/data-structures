package main

import (
	"fmt"
)

// LinkedList is the type that will created a new LinkedList.
type LinkedList struct {
	Root *ListNode
	Size int
}

// ListNode is the Node that will be stored in the LinkedList.
type ListNode struct {
	Value   int
	Message string
	Next    *ListNode
}

// NewLinkedList Returns a new LinkedList.
func NewLinkedList() *LinkedList {
	return &LinkedList{
		Root: nil,
		Size: 0,
	}
}

// AddNode adds a new node to the LinkedList.
func (l *LinkedList) AddNode(value int, message string) *ListNode {
	node := &ListNode{Value: value, Message: message, Next: nil}
	place := l.Root
	if place != nil {
		for place.Next != nil {
			place = place.Next
		}
		place.Next = node
		l.Size++
	} else {
		l.Root = node
		l.Size++
	}
	return node
}

// RemoveNode removes a node from the LinkedList.
func (l *LinkedList) RemoveNode(node *ListNode) {
	place := node.Next
	node.Value = node.Next.Value
	node.Message = node.Next.Message
	node.Next = place.Next

}

// PrintListValues prints out the LinkedList in memory.
func (l *LinkedList) PrintListValues() {
	place := l.Root
	for place != nil {
		println(place.Value)
		place = place.Next
	}
}

// PrintListMessages prints out the LinkedList in memory.
func (l *LinkedList) PrintListMessages() {
	place := l.Root
	for place != nil {
		println(place.Message)
		place = place.Next
	}
}

func main() {
	fmt.Println("Linked Lists!")
	list := NewLinkedList()
	list.AddNode(1, "Hello")
	node1 := list.AddNode(4, "How are you?")
	list.AddNode(5, "How are you?")
	list.PrintListMessages()
	list.RemoveNode(node1)
	list.PrintListMessages()
	fmt.Println(list.Size)
}
