package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
	"sync"
)

// InvertedIndex struct
type InvertedIndex struct {
	index map[string]string
	mutex *sync.Mutex
	size  uint
}

// NewIndex returns a pointer to an Inverted Index object.
func NewIndex() *InvertedIndex {
	// Make a new map that can be given to the InvertedIndex.
	ind := make(map[string]string)
	return &InvertedIndex{
		index: ind,
		mutex: &sync.Mutex{},
		size:  0,
	}
}

// IndexFile reads a file and indexes it.
func (i *InvertedIndex) IndexFile(file string) error {
	// Read the given file into memory. This should be changed in the future.
	fopen, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	// Create a Regex object that can match punctuation within a file's text.
	re, err := regexp.Compile("[.,?]")
	if err != nil {
		return err
	}
	// Remove all punctuation from the file.
	premove := re.ReplaceAllString(string(fopen), "")
	// Split all of the strings within the file.
	str := strings.Split(premove, " ")
	// Now sort all of the words within the file.
	sort.Strings(str)
	// Enter all of the values found in the file into the index.
	i.mutex.Lock()
	for _, word := range str {
		i.index[word] = file
		// Increase the size count of the index.
		i.size++
	}
	i.mutex.Unlock()
	// Return nil if no error occurs.
	return nil
}

// PrintIndex prints the given key's index.
func (i *InvertedIndex) PrintIndex() {
	fmt.Println(i.index)
}

func main() {
	i := NewIndex()
	i.IndexFile("./test.txt")
	i.IndexFile("./test.1.txt")
	i.PrintIndex()
}
