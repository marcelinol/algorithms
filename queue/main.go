package main

import (
	"errors"
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

func main() {}
