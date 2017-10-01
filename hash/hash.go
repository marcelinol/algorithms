package main

import (
	"fmt"
	"strconv"
)

const hashSize int = 30

type HashTable struct {
	// hashFunction func(string) int
	length  int
	storage []string
}

func NewHashTable() *HashTable {
	return &HashTable{0, make([]string, hashSize)}
}

func (h *HashTable) Add(key string, value string) error {
	index := hashFunction(key)
	h.storage[index] = value
	h.length++
	return nil
}

func hashFunction(key string) int {
	value, _ := strconv.Atoi(key)
	return value
}

func main() {
	h := NewHashTable()
	fmt.Println(h)

	fmt.Printf("Adding value 1 to hash\n")
	h.Add("1", "first value")
	fmt.Println(h)

	fmt.Printf("Adding value 23 to hash\n")
	h.Add("23", "xunda")
	fmt.Println(h)
}

// Questions
// Why the length++ didn't work when I was not using a pointer to the hashtable?
