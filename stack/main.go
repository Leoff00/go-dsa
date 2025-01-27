package stack

import "errors"

const max_items_size = 10

type Items struct {
	Item any
	Size int
}

type Stack []Items

var st = make(Stack, 0, max_items_size)

func (s *Stack) IsFull() bool {
	return len(st) == max_items_size
}

func (s *Stack) IsEmpty() bool {
	return len(st) == 0
}

func (s *Stack) Unstack() (int, error) {
	if !st.IsEmpty() {
		st = st[:len(st)-1]
		return len(st), nil
	}
	return 0, errors.New("the Stack is already empty")
}

func (s *Stack) Pile(i Items) (int, error) {
	if !s.IsFull() {
		st = append(st, i)
		return len(st), nil
	}
	return 0, errors.New("the Stack is already full")
}
