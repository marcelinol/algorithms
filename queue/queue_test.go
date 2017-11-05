package main

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()

	result := fmt.Sprintf("%T", q)

	if result != "*main.Queue" {
		t.Error("Expected q to be a Queue but got ", result)
	}

	front := q.front
	if front != 0 {
		t.Error("Expected q.front to be 0 but got ", front)
	}

	rear := q.rear
	if rear != 0 {
		t.Error("Expected q.rear to be 0 but got ", rear)
	}
}

func TestPeek(t *testing.T) {
	q := NewQueue()

	_, err := q.Peek()

	if err == nil {
		t.Error("Expected q.Peek() to return an error but it didn't")
	}

	el := "first"
	q.Add(el)
	peek, err := q.Peek()

	if err != nil {
		t.Error("Test failed, error must be nil")
	}

	if peek != el {
		t.Error("Expected peek to be " + el + ", but got " + peek)
	}

	secondEl := "second"
	q.Add(secondEl)
	peek, err = q.Peek()

	if err != nil {
		t.Error("Test failed, error must be nil")
	}

	if peek != el {
		t.Error("Expected peek to be " + el + ", but got " + peek)
	}
}

func TestFIFO(t *testing.T) {
	q := NewQueue()

	first := "first"
	second := "second"
	err := q.Add(first)

	if err != nil {
		t.Error("Test failed, error must be nil")
	}

	err = q.Add(second)

	if err != nil {
		t.Error("Test failed, error must be nil")
	}

	r, err := q.Remove()

	if err != nil {
		t.Error("Test failed, error must be nil")
	}

	if r != first {
		t.Error("Expected removed item to be first but got ", r)
	}

	r, err = q.Remove()

	if err != nil {
		t.Error("Test failed, error must be nil")
	}

	if r != second {
		t.Error("Expected removed item to be second but got ", r)
	}

	_, err = q.Remove()

	if err == nil {
		t.Error("Expected error to be 'Queue is empty'")
	}
}

func TestResize(t *testing.T) {
	q := NewQueue()

	for n := 0; n <= 5000; n++ {
		el := fmt.Sprintf("element %v", n)
		err := q.Add(el)
		if err != nil {
			t.Error("Expected errors to be nil")
		}
	}
}
