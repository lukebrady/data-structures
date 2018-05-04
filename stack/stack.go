package main

import "sync"
import "fmt"

// Stack is the stack type that will be used to create
// a stack object within the program.
type Stack struct {
	stackArr []int
	size     uint32
	lock     *sync.Mutex
}

// NewStack creates a new stack object.
func NewStack() *Stack {
	sArr := []int{}
	return &Stack{
		stackArr: sArr,
		size:     0,
		lock:     &sync.Mutex{},
	}
}

// Push an item to the stack.
func (s *Stack) push(item int) {
	// Lock the stack so that the item can be put into the stack.
	s.lock.Lock()
	// Start a new go routine to append to the stack and increment the size.
	s.stackArr = append(s.stackArr, item)
	s.size++
	// Unlock the stack so that more operations can be done to the stack.
	s.lock.Unlock()
}

// Pop removes the last inserted item from the stack.
func (s *Stack) pop() {
	// Lock the stack so that the last value can be removed.
	s.lock.Lock()
	// A new stack slice is created.
	if s.size != 0 {
		fmt.Printf("Popping %d from the stack.\n", s.stackArr[s.size-1])
		s.stackArr = s.stackArr[0 : s.size-1]
		s.size--
	} else {
		println("The stack is empty.")
	}
	// Unlock the stack for use by other threads.
	s.lock.Unlock()
}

// Print all contents of the stack.
func (s *Stack) print() {
	for _, item := range s.stackArr {
		println(item)
	}
}

func main() {
	stack := NewStack()
	stack.push(15)
	stack.push(25)
	stack.push(36)
	stack.push(45)
	stack.push(56)
	stack.print()
	stack.pop()
	stack.pop()
	stack.pop()
	stack.pop()
	stack.print()
}
