package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting Linked List")

	list := NewLinkedList()
	fmt.Println(list)

	fmt.Println("Adding rocky balboa to list")
	list.AddFirst("rocky balboa")
	fmt.Println(list)

	fmt.Println("Searching for apolo creed")
	fmt.Println(list.include("apolo creed"))

	fmt.Println("Searching for rocky balboa")
	fmt.Println(list.include("rocky balboa"))
}

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
