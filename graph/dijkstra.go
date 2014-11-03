package graph

import (
	"errors"
)

func (g *graph) Path(source string, target string) (path []string, err error) {
	//priority queue
	pq := make(priority, 0)
	//distance from source to specified node
	dist := make(map[*node]int)
	//whether this node has been visited
	visited := make(map[*node]bool)
	//node visited previous to this one (for reconstruction of path)
	previous := make(map[*node]*node)

	//Choose best option from priority queue
	best, ok := g.nodes[source]
	//If source isn't in graph, error
	if !ok {
		path = nil
		err = errors.New("Source not in graph.")
		return
	}
	for len(pq) > 0 {

		//Is it the target?
		if best.value == target {
			break
		}
		//For each neighbor of best, check if visited
		for v, w := range best.adjacent {
			if !visited[v] {
				altDist := dist[best] + w
				curDist, ok := dist[v]
				if !ok || altDist < curDist {
					dist[v] = altDist
					pq.Add(v, altDist)
					previous[v] = best
				}
			}
			//Pull next node
			best = pq.Next()
		}
	}
	//Did we find target?
	if best == nil {
		path = make([]string, 0)
		err = nil
		return
	}
	//Reconstruct path
	path = reverse(reconstruct(previous, best, source))
	err = nil
	return
}

func reconstruct(prev map[*node]*node, current *node, target string) []string {
	ans := make([]string, 0)
	for current.value != target {
		ans = append(ans, current.value)
		current = prev[current]
	}
	ans = append(ans, current.value)
	return ans
}

func reverse(a []string) []string {
	i := len(a) - 1
	j := 0
	for j < i {
		a[i], a[j] = a[j], a[i]
		i++
		j++
	}
	return a
}
