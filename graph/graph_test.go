package graph

import (
	"search"
	"testing"
)

func GraphForTest() *graph {
	g := New()
	a := "a"
	b := "b"
	c := "c"
	d := "d"
	e := "e"
	g.Add(a)
	g.Add(b)
	g.Add(c)
	g.Add(d)
	g.Add(e)
	g.AddEdge(a, b)
	g.AddEdge(b, c)
	g.AddEdge(c, d)
	g.AddEdge(d, b)
	g.AddEdge(d, e)
	return &g
}

func TestAdjacent(t *testing.T) {
	g := GraphForTest()
	s := g.AdjacentTo(b)
	if len(s) != 3 {
		t.Error("Too few/many adjacent nodes found")
	}
	contained := [...]string{a, c, d}
	s = search.Strings(s)
	for i, v := range contained {
		if s[i] != v {
			t.Errorf("%s should be %s", s[i], v)
		}
	}
}

func TestVertices(t *testing.T) {
	g := GraphForTest()
	if g.Vertices() != 5 {
		t.Errorf("There should be %d vertices (found %d)", 5, g.Vertices())
	}
}

func TestEdges(t *testing.T) {
	g := GraphForTest()
	if g.Edges() != 5 {
		t.Errorf("There should be %d edges (found %d)", 5, g.Edges())
	}
}

func TestZeroEdges(t *testing.T) {
	g := New()
	g.Add("a")
	g.Add("b")
	if g.Edges() != 0 {
		t.Errorf("There should be %d edges (found %d)", 0, g.Edges())
	}
}

func TestZeroVertices(t *testing.T) {
	g := New()
	if g.Vertices() != 0 {
		t.Errorf("There should be %d vertices (found %d)", 0, g.Vertices())
	}
}

func TestZeroAdjacent(t *testing.T) {
	g := New()
	f := "f"
	g.Add(f)
	s := g.AdjacentTo(f)
	if len(s) != 0 {
		t.Errorf("There should be %d zero adjacent (found %d)", 0, len(s))
	}
}

//TODO: Test String()?
