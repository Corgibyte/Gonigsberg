//Package digraph provides the interface and implementation of a directed graph. Note: because of the lack of Go generics, this implementation does not have guaranteed compile-time type safety.

package digraph

type Interface interface {
	New() Digraph
	Vertices() int
	Edges() int
	Add(x interface{})
	ConnectedTo(x interface{}) []interface{}
	Reverse() Digraph
	String() string
}
