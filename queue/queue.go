package main

import "sync"
import "fmt"

// Queue is the queue type that will be used to create
// a Queue object within the program.
type Queue struct {
	QueueArr []int
	size     uint32
	pos      uint32
	lock     *sync.Mutex
}

// NewQueue creates a new queue object.
func NewQueue() *Queue {
	sArr := []int{}
	return &Queue{
		QueueArr: sArr,
		size:     0,
		pos:      0,
		lock:     &sync.Mutex{},
	}
}

// enqueue an item to the Queue.
func (s *Queue) enqueue(item int) {
	// Lock the Queue so that the item can be put into the Queue.
	s.lock.Lock()
	// Start a new go routine to append to the Queue and increment the size.
	s.QueueArr = append(s.QueueArr, item)
	s.size++
	// Unlock the Queue so that more operations can be done to the Queue.
	s.lock.Unlock()
}

// dequeue removes the last inserted item from the Queue.
func (s *Queue) dequeue() {
	// Lock the Queue so that the last value can be removed.
	s.lock.Lock()
	// A new Queue slice is created.
	if s.size != 0 {
		fmt.Printf("dequeueping %d from the Queue.\n", s.QueueArr[s.pos])
		s.QueueArr = s.QueueArr[s.pos+1:]
		s.size--
	} else {
		println("The Queue is empty.")
	}
	// Unlock the Queue for use by other threads.
	s.lock.Unlock()
}

// Print all contents of the Queue.
func (s *Queue) print() {
	for _, item := range s.QueueArr {
		println(item)
	}
}

func main() {
	Queue := NewQueue()
	Queue.enqueue(15)
	Queue.enqueue(25)
	Queue.enqueue(36)
	Queue.enqueue(45)
	Queue.enqueue(56)
	Queue.print()
	Queue.dequeue()
	Queue.dequeue()
	Queue.dequeue()
	Queue.dequeue()
	Queue.dequeue()
	Queue.dequeue()
	Queue.print()
}
