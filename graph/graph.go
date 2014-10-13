//Package graph provides the interface and implementation of a basic, undirected graph. Note: because of the lack of Go generics, this implementation does not have guaranteed compile-time type safety.

package graph

type Interface interface {
	New() Graph
	Vertices() int
	Edges() int
	Add(x interface{})
	AdjacentTo(x interface{}) []interface{}
	String() string
}
