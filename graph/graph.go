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
type Graph struct {
	nodes    map[string]*node
	vertices int
	edges    int
}

//Get new, blank graph
func New() *Graph {
	return &Graph{make(map[string]*node), 0, 0}
}

//Number of vertices in the graph
func (g *Graph) Vertices() int {
	return g.vertices
}

//Number of edges in the Graph
func (g *Graph) Edges() int {
	return g.edges
}

//Add vertice to the Graph
func (g *Graph) Add(x string) (err error) {
	_, ok := g.nodes[x]
	if ok {
		err = errors.New(fmt.Sprintf("%v already in Graph", x))
	}
	n := node{make(map[*node]int), x}
	g.nodes[x] = &n
	g.vertices++
	return nil
}

//Add an edge between two specified members of the graph
//Returns an error if a or b are not in graph
//Returns an error if weight is negative
func (g *Graph) AddEdge(a string, b string, weight int) (err error) {
	if weight < 0 {
		err = errors.New(fmt.Sprintf("Cannot have negative weight: %d", weight))
		return
	}
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
//Returns an error if s is not in the Graph
func (g *Graph) AdjacentTo(s string) ([]string, error) {
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
func (g *Graph) AdjacentEdges(s string) (map[string]int, error) {
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
