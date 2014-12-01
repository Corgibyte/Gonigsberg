package Gonigsberg

import (
	"bufio"
	"github.com/Corgibyte/Gonigsberg/graph"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	// Create graph and list of locations
	graph := graph.New()

	// Add locations to graph
	locs := addLocationsToGraph(bufio.NewScanner(os.Open("data/locations.txt")), graph)

	// Add edges to graph

}

func addLocationsToGraph(sc *bufio.Scanner, g *graph.Graph) *[]location {
	locs := make(*[]location, 0)
	gen := make(serialGenerator, 0)
	for sc.Scan() {
		l := NewLocation(gen.generateSerial(), sc.Text())
		locs = append(locs, l)
		g.add(l)
	}
	return locs
}

func addEdges(sc *bufio.Scanner, g *graph.Graph) {
	

type serialGenerator map[string]bool

func (s *serialGenerator) generateSerial() string {
	random := strconv.FormatInt(rand.Int63())
	for s[random] {
		random = strconv.FormatInt(rand.Int63())
	}
	s[random] = true
	return random
}
