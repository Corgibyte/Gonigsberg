package graph

import (
	"testing"
)

func BasicGraph() *graph {
	g := New()
	g.Add("a")
	g.Add("b")
	g.Add("c")
	g.Add("d")
	g.Add("e")
	g.Add("z")
	g.AddEdge("a", "b", 1)
	g.AddEdge("b", "c", 2)
	g.AddEdge("c", "z", 3)
	g.AddEdge("a", "d", 4)
	g.AddEdge("d", "e", 1)
	g.AddEdge("e", "z", 1)
	return g
}

func TestPath(t *testing.T) {
	g := BasicGraph()
	s, _ := g.Path("a", "z")
	expected := [...]string{"a", "b", "c", "z"}
	if SamePath(expected[0:4], s) {
		t.Error("Path not correct.")
	}
}

func TestPathShorterPathHasMoreNodes(t *testing.T) {
	g := New()
	g.Add("a")
	g.Add("b")
	g.Add("c")
	g.Add("d")
	g.Add("z")
	g.AddEdge("a", "b", 1)
	g.AddEdge("b", "c", 1)
	g.AddEdge("c", "z", 1)
	g.AddEdge("a", "d", 1)
	g.AddEdge("d", "z", 5)
	s, _ := g.Path("a", "z")
	expected := [...]string{"a", "b", "c", "z"}
	if SamePath(expected[0:4], s) {
		t.Errorf("Path not correct.")
	}
}

func SamePath(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestNoPath(t *testing.T) {
	g := New()
	g.Add("a")
	g.Add("z")
	s, ok := g.Path("a", "z")
	if s != nil {
		t.Error("Path should be nil!")
	}
	if ok == nil {
		t.Error("Error from empty path shouldn't be nil!")
	}
}

func TestDirectPath(t *testing.T) {
	g := New()
	g.Add("a")
	g.Add("b")
	g.Add("z")
	g.AddEdge("a", "z", 1)
	s, _ := g.Path("a", "z")
	expected := [...]string{"a", "z"}
	if SamePath(expected[0:1], s) {
		t.Error("Path not correct.")
	}
}

func TestChainingPath(t *testing.T) {
	g := BasicGraph()
	g.Add("f")
	g.AddEdge("b", "f", 4)
	s, _ := g.Path("a", "z")
	expected := [...]string{"a", "b", "c", "z"}
	if SamePath(expected[0:4], s) {
		t.Error("Path not correct. Is/Expected:\n%d\n%d")
	}
}
