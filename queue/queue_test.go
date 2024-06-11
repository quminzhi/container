package queue

import (
	"testing"
)

func TestIntQueue_Dequeue(t *testing.T) {
	iq := NewIntQueue()
	iq.Enqueue(1)
	iq.Enqueue(2)
	if iq.Dequeue().(int) != 1 {
		t.Errorf("expected 1, got %d", iq.Dequeue().(int))
	}
	if iq.Dequeue().(int) != 2 {
		t.Errorf("expected 2, got %d", iq.Dequeue().(int))
	}
	if iq.Dequeue() != nil {
		t.Errorf("expected nil, got %v", iq.Dequeue())
	}
}

func TestIntQueue_End(t *testing.T) {
	iq := NewIntQueue()
	if iq.End() != nil {
		t.Errorf("expected nil, got %v", iq.End())
	}
	iq.Enqueue(1)
	iq.Enqueue(2)
	iq.Enqueue(3)
	if iq.End().(int) != 3 {
		t.Errorf("expected 3, got %d", iq.End().(int))
	}
}

func TestIntQueue_Enqueue(t *testing.T) {
	iq := NewIntQueue()
	iq.Enqueue(1)
	if iq.End().(int) != 1 {
		t.Errorf("expected 1, got %d", iq.End().(int))
	}
	iq.Enqueue(1)
	if iq.End().(int) != 1 {
		t.Errorf("expected 1, got %d", iq.End().(int))
	}
	iq.Enqueue(-1)
	if iq.End().(int) != -1 {
		t.Errorf("expected -1, got %d", iq.End().(int))
	}
}

func TestIntQueue_Front(t *testing.T) {
	iq := NewIntQueue()
	if iq.Front() != nil {
		t.Errorf("expected nil, got %v", iq.Front())
	}
	iq.Enqueue(1)
	if iq.Front().(int) != 1 {
		t.Errorf("expected 1, got %d", iq.Front().(int))
	}
}

func TestNewIntQueue(t *testing.T) {
	iq := NewIntQueue()
	if len(iq) != 0 {
		t.Errorf("expected 0, got %d", len(iq))
	}
	if iq.End() != nil {
		t.Errorf("expected nil, got %v", iq.End())
	}
	if iq.Dequeue() != nil {
		t.Errorf("expected nil, got %v", iq.Dequeue())
	}
	if iq.Front() != nil {
		t.Errorf("expected nil, got %v", iq.Front())
	}
}
