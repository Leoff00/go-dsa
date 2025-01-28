package queue

import "errors"

const MAX_SIZE_QUEUE = 10

type Queue []any

func (q *Queue) IsFull() bool {
	return len(*q) == MAX_SIZE_QUEUE
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Pop() (int, error) {
	if !q.IsEmpty() {
		*q = (*q)[1:]
		return len(*q), nil
	}
	return 0, errors.New("queue is empty, can not remove any element")
}

func (q *Queue) Push(value any) (int, error) {
	if !q.IsFull() {
		*q = append(*q, value)
		return len(*q), nil
	}
	return 0, errors.New("queue is full, can not add any element")

}
