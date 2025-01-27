package stack

import (
	"testing"
)

func TestIsFull(t *testing.T) {
	st := make(Stack, 0, max_items_size)

	for i := 0; i < max_items_size; i++ {
		st = append(st, Items{Item: i, Size: i})
	}

	if st.IsFull() {
		t.Log("Expected stack got full")
	}

}

func TestIsEmpty(t *testing.T) {
	//initializing with no values
	st := make(Stack, 0, max_items_size)
	if st.IsEmpty() {
		t.Log("The stack is fresh, with no items")
	}
}

func TestIsNotEmpty(t *testing.T) {
	//initializing with no values
	st := make(Stack, 0, max_items_size)
	item := Items{Item: "test", Size: 1}
	currLen, _ := st.Pile(item)

	if !st.IsEmpty() {
		t.Logf("The stack isn't empty no more, with len %d", currLen)
	}

}

func TestPile(t *testing.T) {
	expected := 1
	st := make(Stack, 0, max_items_size)
	item := Items{Item: "test", Size: 1}
	currLen, _ := st.Pile(item)

	if currLen != expected {
		t.Errorf("Expected stack size 1, but got %d", len(st))
	}
}

func TestPileErr(t *testing.T) {
	st := make(Stack, 0, max_items_size)
	item := Items{Item: "test", Size: 1}
	_, err := st.Pile(item)

	for i := 0; i < max_items_size; i++ {
		st = append(st, Items{Item: i, Size: i})
	}

	if err != nil {
		t.Errorf("The stack is full!")
	}
}

func TestUnstack(t *testing.T) {
	expected := 9
	st := make(Stack, 0, max_items_size)
	for i := 0; i < max_items_size; i++ {
		st = append(st, Items{Item: i, Size: i})
	}

	currLen, _ := st.Unstack()
	if currLen == expected {
		t.Log("The stack now have 9 items")
	}

}
