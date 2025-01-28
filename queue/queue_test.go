package queue

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	expected := 1
	var q Queue

	currLen, _ := q.Push("foo")

	fmt.Println(q)
	if currLen != expected {
		t.Error("The current len must be 1")
	}
}

func TestPop(t *testing.T) {
	var q Queue

	q.Push("foo")

	expectedLength := 0
	currLen, err := q.Pop()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if currLen != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, currLen)
	}

	_, err = q.Pop()
	if err == nil {
		t.Errorf("Expected an error when popping from an empty queue")
	}
}

func TestIsFull(t *testing.T) {
	var q Queue

	for i := 0; i < MAX_SIZE_QUEUE; i++ {
		q.Push(i)
	}

	if !q.IsFull() {
		t.Errorf("Expected queue to be full")
	}

	_, err := q.Push("overflow")
	if err == nil {
		t.Errorf("Expected an error when adding to a full queue")
	}
}

func TestIsEmpty(t *testing.T) {
	var q Queue
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}

	q.Push("foo")
	if q.IsEmpty() {
		t.Errorf("Expected queue to not be empty")
	}
}
