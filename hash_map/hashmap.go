package main

import (
	"crypto/md5"
)

// HashMap is the HashTable type
type HashMap struct {
	hashArray []int
}

// NewHashMap creates a pointer to the HashTable struct.
func NewHashMap() *HashMap {
	var arr []int
	// Create a pointer to a HashMap object and then return.
	return &HashMap{
		hashArray: arr,
	}
}

// Put is a function used to put an item into the hash map.
func (hashMap *HashMap) Put(item string) {
	hash := md5.New()
	hash.Sum([]byte(item))
}

func main() {

}
