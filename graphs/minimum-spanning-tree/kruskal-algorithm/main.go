package main

import (
	"fmt"
	"sort"
)

// pair vertexFrom and vertexTo represents an edge from a vertex vertexFrom to a vertex vertexTo and weight represents the weight of that edge.
type pair struct {
	vertexFrom, vertexTo int
	weight               int
}

// algorithms is adjacency matrix by edge
var graph []pair

// number of vertices and edges in this algorithms
var numberOfVertices, numberOfEdges int

const max = 1e4 + 5

var id [max]int

// pairSorter joins a By function and a slice of pair to be sorted.
type pairSorter struct {
	graph []pair
	by    func(p1, p2 pair) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (p *pairSorter) Len() int {
	return len(p.graph)
}

// Swap is part of sort.Interface.
func (p *pairSorter) Swap(i, j int) {
	p.graph[i], p.graph[j] = p.graph[j], p.graph[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (p *pairSorter) Less(i, j int) bool {
	return p.by(p.graph[i], p.graph[j])
}

func main() {

	// initialize id
	for i := 0; i < max; i++ {
		id[i] = i
	}

	fmt.Scanf("%d", &numberOfVertices)
	fmt.Scanf("%d", &numberOfEdges)

	// adjacency matrix
	graph = make([]pair, numberOfEdges, numberOfEdges)

	// generates the adjacency matrix, the indexes are the number of the edge
	var vertexFrom, vertexTo, weight int
	for i := 0; i < numberOfEdges; i++ {
		fmt.Scanf("%d", &vertexFrom)
		fmt.Scanf("%d", &vertexTo)
		fmt.Scanf("%d", &weight)
		pair := pair{
			vertexFrom: vertexFrom,
			vertexTo:   vertexTo,
			weight:     weight,
		}

		// update data of the edge
		graph[i] = pair
	}

	// sort the edges ascending by weight
	// Closures that order the pair structure.
	weightClosure := func(p1, p2 pair) bool {
		return p1.weight <= p2.weight
	}

	pairSorter := &pairSorter{
		graph: graph,
		by:    weightClosure,
	}

	sort.Sort(pairSorter)
	fmt.Println(graph[0], graph[1])
	minimumCost := kruskal(graph)

	fmt.Println(minimumCost)

}

func kruskal(graph []pair) int {
	var weight, minimumCost, vertexFrom, vertexTo int

	for i := 0; i < numberOfEdges; i++ {
		// Selecting edges one by one in increasing order from the beginning
		vertexFrom = graph[i].vertexFrom
		vertexTo = graph[i].vertexTo
		weight = graph[i].weight

		if root(vertexFrom) != root(vertexTo) {
			minimumCost += weight
			union(vertexFrom, vertexTo)
		}
	}
	return minimumCost
}

func root(x int) int {
	for id[x] != x {

		id[x] = id[id[x]]
		x = id[x]

	}
	return x
}

func union(vertexFrom, vertexTo int) {
	rootVertexFrom := root(vertexFrom)
	rootVertexTo := root(vertexTo)
	id[rootVertexFrom] = id[rootVertexTo]

}
