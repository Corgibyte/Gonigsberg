package graph

import (
	"container/heap"
	"testing"
)

func TestPriority(t *testing.T) {
	h := make(priority, 0)
	h.Add(&node{nil, "a"}, 3)
	h.Add(&node{nil, "b"}, 2)
	h.Add(&node{nil, "c"}, 4)
	heap.Init(&h)
	x := heap.Pop(&h).(*edge)
	if x.target.value != "b" {
		t.Errorf("Wrong target found: %v", x.target.value)
	}
}

func TestPopEmpty(t *testing.T) {
	h := make(priority, 0)
	h.Add(&node{nil, "a"}, 1)
	heap.Init(&h)
	x := heap.Pop(&h).(*edge)
	if x.target.value != "a" {
		t.Errorf("Something wrong with the priority queue.")
	}
}
