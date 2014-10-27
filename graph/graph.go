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

func (n *node) Value() string {
	return n.value
}

//Ths undirected graph holds string and is a mutable graph with edge weight 1
type graph struct {
	nodes    map[string]*node
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
	_, ok := g.nodes[x]
	if ok {
		err = errors.New(fmt.Sprintf("%v already in graph", x))
	}
	n := node{make(map[*node]int), x}
	g.nodes[x] = &n
	g.vertices++
	return nil
}

//Add an edge between two specified members of the graph
//Returns an error if a or b are not in graph
func (g *graph) AddEdge(a string, b string, weight int) (err error) {
	aNode, aInGraph := g.nodes[a]
	bNode, bInGraph := g.nodes[b]
	if !aInGraph {
		err = errors.New(fmt.Sprintf("%v not found in graph", a))
		return
	} else if !bInGraph {
		err = errors.New(fmt.Sprintf("%v not found in graph", b))
		return
	}
	aNode.adjacent[bNode] = weight
	bNode.adjacent[aNode] = weight
	g.edges++
	return nil
}

//Gets all values that are adjacent (directly linked) to the specified value
//Returns an error if s is not in the graph
func (g *graph) AdjacentTo(s string) ([]string, error) {
	sNode, ok := g.nodes[s]
	if !ok {
		err := errors.New(fmt.Sprintf("%v not in graph", s))
		return nil, err
	}
	adjStrings := make([]string, 0)
	for adj, _ := range sNode.adjacent {
		adjStrings = append(adjStrings, adj.value)
	}
	return adjStrings, nil
}

//Gets all values that are adjacent to the specified value and the weight
//Of the edges between them
func (g *graph) AdjacentEdges(s string) (map[string]int, error) {
	sNode, ok := g.nodes[s]
	if !ok {
		err := errors.New(fmt.Sprintf("%v not in graph", s))
		return nil, err
	}
	adjEdges := make(map[string]int)
	for n, weight := range sNode.adjacent {
		adjEdges[n.Value()] = weight
	}
	return adjEdges, nil
}
