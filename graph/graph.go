//Package graph provides the interface and implementation of a basic, undirected graph. Note: because of the lack of Go generics, this implementation does not have guaranteed compile-time type safety.

package graph

import (
	"errors"
	"fmt"
)

//A node has a map of nodes it's connected to and the weight between them
type node struct {
	adjacent map[*node]int
	value    string
}

//Ths undirected graph holds string and is a mutable graph with edge weight 1
type graph struct {
	m        map[string]*node
	vertices int
	edges    int
}

//Get new, blank graph
func New() *graph {
	return &graph{make(map[string]*node), 0, 0}
}

//Number of vertices in the graph
func (g *graph) Vertices() int {
	return g.vertices
}

//Number of edges in the graph
func (g *graph) Edges() int {
	return g.edges
}

//Add vertice to the graph
func (g *graph) Add(x string) (err error) {
	_, ok := g.m[x]
	if ok {
		err = errors.New(fmt.Sprintf("%v already in graph", x))
	}
	n := node{make(map[*node]int), x}
	g.m[x] = &n
	g.vertices++
	return nil
}

func (g *graph) AddEdge(a string, b string) (err error) {
	aNode, aInGraph := g.m[a]
	bNode, bInGraph := g.m[b]
	if !aInGraph {
		err = errors.New(fmt.Sprintf("%v not found in graph", a))
		return
	} else if !bInGraph {
		err = errors.New(fmt.Sprintf("%v not found in graph", b))
		return
	}
	aNode.adjacent[bNode] = 1
	bNode.adjacent[aNode] = 1
	return nil
}

func (g *graph) AdjacentTo(s string) ([]string, error) {
	sNode, ok := g.m[s]
	if !ok {
		err := errors.New(fmt.Sprintf("%v not in graph", s))
		return nil, err
	}
	adjStrings := make([]string, 5)
	for adj, _ := range sNode.adjacent {
		adjStrings = append(adjStrings, adj.value)
	}
	return adjStrings, nil
}
