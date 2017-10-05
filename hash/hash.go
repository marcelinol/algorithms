package main

import (
	"fmt"
	"strconv"
)

// Linked List

type Node struct {
	item     interface{}
	nextNode *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func NewLinkedList() *LinkedList {
	node := Node{nil, nil}

	return &LinkedList{&node, 0}
}

func (l *LinkedList) AddFirst(item interface{}) {
	newNode := Node{item, l.head}
	l.head = &newNode
	l.length++
}

func (l *LinkedList) include(item interface{}) bool {
	currentNode := l.head
	for currentNode.nextNode != nil {
		if currentNode.item == item {
			return true
		}
		currentNode = currentNode.nextNode
	}

	return false
}

// HashTable

const hashSize int = 30

type HashItem struct {
	key   string
	value string
}

type HashTable struct {
	length  int
	storage [hashSize]LinkedList
}

func NewHashTable() *HashTable {
	return &HashTable{0, [hashSize]LinkedList{}}
}

func (h *HashTable) Add(key string, value string) {
	index := hashFunction(key)

	fmt.Printf("Adding value " + value + " to index ")
	fmt.Println(index)
	fmt.Println(h.storage[index])

	newItem := HashItem{key, value}

	h.storage[index].AddFirst(newItem)
	h.length++
}

func (h *HashTable) Search(key string) string {
	index := hashFunction(key)

	currentNode := h.storage[index].head

	for currentNode.nextNode != nil {
		if currentNode.item.key == key {
			return currentNode.item.value
		}
		currentNode = currentNode.nextNode
	}

	item := currentNode.item
	return item.value
}

func hashFunction(key string) int {
	value, _ := strconv.Atoi(key)
	return value % 30
}

func main() {
	h := NewHashTable()
	// fmt.Println(h)

	h.Add("1", "first value")
	// fmt.Println(h)

	h.Add("20", "xunda")
	// fmt.Println(h)

	h.Add("50", "cinquentao") // It should cause a collision with the "xunda" item
	// fmt.Println(h)

	h.Add("20", "this should overwrite the item xunda")
	fmt.Println(h)

	vintao := h.Search("20")
	fmt.Println(vintao)

	cinquentao := h.Search("50")
	fmt.Println(cinquentao)
}

// Questions
// Why the length++ didn't work when I was not using a pointer to the hashtable?
