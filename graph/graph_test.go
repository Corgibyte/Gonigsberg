package graph

import (
	"sort"
	"testing"
)

//Creates graph used for some testing
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
	g.AddEdge(a, b, 1)
	g.AddEdge(b, c, 1)
	g.AddEdge(c, d, 1)
	g.AddEdge(d, b, 1)
	g.AddEdge(d, e, 1)
	return g
}

func TestAdjacent(t *testing.T) {
	g := GraphForTest()
	s, ok := g.AdjacentTo("b")
	if ok != nil {
		t.Logf("Adjacent doesn't work, message: %v", ok)
	}
	if len(s) != 3 {
		t.Errorf("Too few/many adjacent nodes found: %d should be 3", len(s))
	}
	contained := [...]string{"a", "c", "d"}
	sort.Strings(s)
	for i, v := range contained {
		if s[i] != v {
			t.Errorf("%s should be %s", s[i], v)
		}
	}
}

func TestEdgeWeights(t *testing.T) {
	g := New()
	g.Add("a")
	g.Add("b")
	g.Add("c")
	g.AddEdge("a", "b", 2)
	g.AddEdge("a", "c", 4)
	edges, _ := g.AdjacentEdges("a")
	if len(edges) != 2 {
		t.Error("Wrong number of edges returned")
	}
	bWeight, bOk := edges["b"]
	cWeight, cOk := edges["c"]
	if !bOk || !cOk {
		t.Error("Didn't return one of the edges.")
	}
	if bWeight != 2 {
		t.Errorf("bWeight wrong: %d", bWeight)
	}
	if cWeight != 4 {
		t.Errorf("cWeight wrong: %d", cWeight)
	}
}

func TestZeroEdgeWeight(t *testing.T) {
	g := New()
	g.Add("a")
	g.Add("b")
	g.AddEdge("a", "b", 0)
	edges, _ := g.AdjacentEdges("a")
	if len(edges) != 1 {
		t.Error("Wrong number of edges returned")
	}
	weight, _ := edges["b"]
	if weight != 0 {
		t.Errorf("Weight wrong: %d", weight)
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
	s, ok := g.AdjacentTo(f)
	if ok != nil {
		t.Logf("Adjacent doesn't work, message: %v", ok)
	}
	if len(s) != 0 {
		t.Errorf("There should be %d zero adjacent (found %d)", 0, len(s))
	}
}
