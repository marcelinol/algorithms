package main

import (
	"fmt"
	"strconv"
)

// Linked List

type node struct {
	key   string
	value string
	next  *node
}

type linkedList struct {
	first node
	len   int
}

func (n node) getKey() string {
	return n.key
}

func (n node) getValue() string {
	return n.value
}

func (n *node) setKey(key string) {
	n.key = key
}

func (n *node) setValue(value string) {
	n.value = value
}

func (n node) getNext() node {
	if n.next == nil {
		return node{}
	} else {
		return *n.next
	}
}

func (n *node) setNext(value node) {
	n.next = &value
}

func (list *linkedList) add(key string, value string, index int) {

	if index >= 0 {
		list_node := node{key, value, nil}

		if list.empty() {
			list.first = list_node
		} else {
			if index == 0 {
				// insert in first
				list_node.setNext(list.first)
				list.first = list_node
			} else {
				prev_node := list.first
				curr_node := list.first.getNext()
				curr_index := 1

				for curr_node.next != nil {

					if curr_index == index {
						list_node.setNext(curr_node)
						prev_node.setNext(list_node)
						break
					}

					prev_node = curr_node
					curr_node = curr_node.getNext()
					curr_index++
				}
			}
		}
		list.len++
	}
}

func (list linkedList) remove(index int) {

	if !list.empty() && index >= 0 && index < list.len {
		flag_remove := false

		if list.first.next == nil {
			list.first = node{}
			flag_remove = true
		} else if index == 0 {
			list.first = list.first.getNext()
			flag_remove = true
		} else {
			prev_node := list.first
			curr_node := list.first.getNext()
			curr_index := 1

			for curr_node.next != nil {
				if index == curr_index {
					next_node := curr_node.getNext()
					prev_node.setNext(next_node)
					curr_node.setNext(node{})
					flag_remove = true
					break
				}

				prev_node = curr_node
				curr_node = curr_node.getNext()
				curr_index++
			}
		}

		if flag_remove == true {
			list.len--
		}
	}
}

func (list linkedList) empty() bool {
	if list.len == 0 {
		return true
	} else {
		return false
	}
}

func (list linkedList) show() {
	curr_node := list.first

	for curr_node.next != nil {
		fmt.Println(curr_node.getKey())
		fmt.Println(curr_node.getValue())
		curr_node = curr_node.getNext()
		fmt.Println("Final:", curr_node)
	}
}

// HashTable

const hashSize int = 30

type HashTable struct {
	length  int
	storage [hashSize]linkedList
}

func NewHashTable() *HashTable {
	// Why can't I create the strings array like this?
	// [hashSize]string{}
	return &HashTable{0, [hashSize]linkedList{}}
}

func (h *HashTable) Add(key string, value string) error {
	index := hashFunction(key)

	fmt.Printf("Adding value " + value + " to index ")
	fmt.Println(index)
	fmt.Println(h.storage[index])
	h.storage[index].add(key, value, 0)
	h.length++
	return nil
}

func hashFunction(key string) int {
	value, _ := strconv.Atoi(key)
	return value % 31
}

func main() {
	h := NewHashTable()
	// fmt.Println(h)

	h.Add("1", "first value")
	// fmt.Println(h)

	h.Add("20", "xunda")
	// fmt.Println(h)

	h.Add("50", "xunda2") // It should cause a collision with the "xunda" item
	// fmt.Println(h)

	h.Add("20", "this should overwrite the item xunda")
	fmt.Println(h)
}

// Questions
// Why the length++ didn't work when I was not using a pointer to the hashtable?
