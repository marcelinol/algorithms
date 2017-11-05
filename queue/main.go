package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	storage []string
	front   int
	rear    int
}

const queueSize = 5

func NewQueue() *Queue {
	return &Queue{make([]string, queueSize), 0, 0}
}

func (q *Queue) IsEmpty() bool {
	return q.front == q.rear
}

func (q *Queue) Peek() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("Queue is empty")
	}

	return q.storage[q.front], nil
}

func (q *Queue) Add(s string) error {
	if q.rear == len(q.storage) {
		q.resize()
	}

	q.storage[q.rear] = s
	q.rear++
	return nil
}

func (q *Queue) Remove() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("Queue is empty")
	}

	el, err := q.Peek()
	if err != nil {
		return "", err
	}
	q.front++

	return el, nil
}

func (q *Queue) resize() {
	copy := q.storage
	q.storage = make([]string, len(q.storage)*2)
	q.rear = 0
	q.front = 0

	for i := 0; i < len(copy); i++ {
		q.Add(copy[i])
	}

	return
}

func main() {
	q := NewQueue()

	// // Checking peek error when queue is empty
	// _, err := q.Peek()
	// fmt.Println(err)
	//
	// // Setup
	// _ = q.Add("first")
	// _ = q.Add("second")
	//
	// // Testing Peek
	// peek, _ := q.Peek()
	// fmt.Println(peek)
	//
	// // Testing Remove
	// removed, _ := q.Remove()
	// fmt.Println(removed)
	//
	// removed, _ = q.Remove()
	// fmt.Println(removed)
	//
	// _, err = q.Remove()
	// fmt.Println(err)

	// Lets NOT break this shit
	for n := 0; n <= 5000; n++ {
		el := fmt.Sprintf("element %v", n)
		err := q.Add(el)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("added " + el)
		}
	}
}
