package stack

import (
	"testing"
)

func TestIntStack_Init(t *testing.T) {
	is := Init()
	if len(is) != 0 {
		t.Errorf("len(is) = %d, want 0", len(is))
	}
}

func TestIntStack_Peek(t *testing.T) {
	is := Init()
	is.Push(1)
	if is.Peek() != 1 {
		t.Errorf("Peek() = %d, want 1", is.Peek())
	}
}

func TestIntStack_Push(t *testing.T) {
	is := Init()
	is.Push(1)
	if is.Peek().(int) != 1 {
		t.Errorf("Peek() = %d, want 1", is.Peek().(int))
	}
	is.Push(2)
	if is.Peek().(int) != 2 {
		t.Errorf("Peek() = %d, want 2", is.Peek().(int))
	}
	is.Push(3)
	if is.Peek().(int) != 3 {
		t.Errorf("Peek() = %d, want 3", is.Peek().(int))
	}
	is.Push(4)
	if is.Peek().(int) != 4 {
		t.Errorf("Peek() = %d, want 4", is.Peek().(int))
	}
	if is.Len() != 4 {
		t.Errorf("Len() = %d, want 4", is.Len())
	}
}

func TestIntStack_Pop(t *testing.T) {
	is := Init()
	is.Push(1)
	is.Push(2)
	is.Push(3)
	is.Push(4)
	if x := is.Pop().(int); x != 4 {
		t.Errorf("Pop() = %d, want 4", x)
	}
	if x := is.Pop().(int); x != 3 {
		t.Errorf("Pop() = %d, want 3", x)
	}
	if is.Len() != 2 {
		t.Errorf("Len() = %d, want 2", is.Len())
	}
	if x := is.Pop().(int); x != 2 {
		t.Errorf("Pop() = %d, want 2", x)
	}
	if x := is.Pop().(int); x != 1 {
		t.Errorf("Pop() = %d, want 1", x)
	}
	if is.Len() != 0 {
		t.Errorf("Len() = %d, want 0", is.Len())
	}
}
