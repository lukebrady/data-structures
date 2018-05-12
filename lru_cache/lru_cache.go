package main

import (
	"fmt"
)

// LRUCache struct
type LRUCache struct {
	cache    map[int]int
	queue    []int
	capacity uint
	size     uint
}

// NewLRUCache returns a pointer to an cache object.
func NewLRUCache(capacity uint) *LRUCache {
	cache := make(map[int]int)
	queue := []int{}
	return &LRUCache{
		cache:    cache,
		queue:    queue,
		capacity: capacity,
		size:     0,
	}
}

// Get retrieves a value from the cache.
func (l *LRUCache) Get(key int) int {
	// Check the map to see if the key exists.
	if _, exists := l.cache[key]; exists {
		return l.cache[key]
	}
	// If the key does not exist, return -1 (Cache miss).
	return -1
}

// Put inserts a value into the cache.
func (l *LRUCache) Put(key, value int) uint {
	// Check the capacity of the cache. If the cache is not full
	// append the new value to the cache.
	if l.size < l.capacity {
		// Check to see if the key exists within the cache. If so, return 1.
		if _, exists := l.cache[key]; exists {
			return 1
		}
		// Now add the key and value to the map.
		l.cache[key] = value
		// Append the new value to the queue.
		l.queue = append(l.queue, key)
		// And increase the total size of the queue.
		l.size++
	} else {
		// Remove the least recently used item from the map and queue.
		delete(l.cache, l.queue[0])
		// Remove the LRU item from the queue.
		l.queue = l.queue[1:]
		// Decrement the total size of the queue.
		l.size--
		// Now call the Put function again.
		l.Put(key, value)
	}
	return 0
}

// PrintCache prints the cache to the screen.
func (l *LRUCache) PrintCache() {
	fmt.Println(l.cache)
}

// PrintKey prints the value of a key to the screen.
func (l *LRUCache) PrintKey(key int) {
	fmt.Println(l.Get(key))
}

func main() {
	cache := NewLRUCache(5)
	cache.Put(1, 3)
	cache.Put(4, 2)
	cache.Put(4, 3)
	cache.Put(5, 8)
	cache.Put(6, 9)
	cache.Put(10, 12)
	cache.Put(12, 13)
	cache.Put(14, 3)
	cache.Put(6, 4)
	cache.Put(5, 8)
	cache.Put(7, 3)
	cache.PrintKey(7)
	cache.PrintCache()
}
