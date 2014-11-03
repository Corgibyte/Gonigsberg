package graph

import (
	"container/heap"
)

type edge struct {
	target *node
	weight int
	index  int
}

type priority []*edge

func (p *priority) Next() *node {
	x := heap.Pop(p).(*edge)
	return x.target
}

func (p *priority) Add(a *node, i int) {
	heap.Push(p, &edge{a, i, -1})
	heap.Init(p)
}

func (p *priority) Push(x interface{}) {
	n := len(*p)
	item := x.(*edge)
	item.index = n
	*p = append(*p, item)
}

func (p *priority) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	item.index = -2 //for safety
	*p = old[0 : n-1]
	return item
}

func (p priority) Len() int {
	return len(p)
}

func (p priority) Less(i, j int) bool {
	return p[i].weight < p[j].weight
}

func (p priority) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = j
	p[j].index = i
}
