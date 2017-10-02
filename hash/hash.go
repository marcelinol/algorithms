package main

import (
	"fmt"
	"strconv"
)

const hashSize int = 30

type HashTable struct {
	length  int
	storage []string
}

func NewHashTable() *HashTable {
	// Why can't I create the strings array like this?
	// [hashSize]string{}
	return &HashTable{0, make([]string, hashSize)}
}

func (h *HashTable) Add(key string, value string) error {
	index := hashFunction(key)

	for h.storage[index] != "" {
		fmt.Printf("Collision on index ")
		fmt.Println(index)
		index++
	}

	fmt.Printf("Adding value " + value + " to index ")
	fmt.Println(index)
	h.storage[index] = value
	h.length++
	return nil
}

func hashFunction(key string) int {
	value, _ := strconv.Atoi(key)
	return value % 31
}

func main() {
	h := NewHashTable()
	fmt.Println(h)

	h.Add("1", "first value")
	fmt.Println(h)

	h.Add("20", "xunda")
	fmt.Println(h)

	h.Add("50", "xunda2") // It should cause a collision with the "xunda" item
	fmt.Println(h)

	h.Add("20", "this should overwrite the item xunda")
	fmt.Println(h)

}

// Questions
// Why the length++ didn't work when I was not using a pointer to the hashtable?
