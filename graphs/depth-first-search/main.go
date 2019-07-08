package main

import (
	"fmt"
)

//  slice to mark if the node is visited
var vis []bool
// adjacency matrix
var graph [][]int
// number of nodes reachable
var numberOfNodesReachable int

func main() {

	// number of nodes and edges in this algorithms
	var numberOfNodes, numberOfEdges int

	fmt.Scanf("%d", &numberOfNodes)
	fmt.Scanf("%d", &numberOfEdges)

	// adjacency matrix
	graph = make([][]int, numberOfNodes, numberOfNodes)

	// generates the adjacency matrix
	for i := 0; i < numberOfEdges; i++ {
		var nodeFrom, nodeTo int
		fmt.Scanf("%d", &nodeFrom)
		fmt.Scanf("%d", &nodeTo)
		graph[nodeFrom-1] = append(graph[nodeFrom-1], nodeTo)
		graph[nodeTo-1] = append(graph[nodeTo-1], nodeFrom)
	}

	// The next line consists of a single integer x denoting the index of the head node.
	var headNode int
	fmt.Scanf("%d", &headNode)

	// slice to mark if the node is visited
	vis = make([]bool, numberOfNodes, numberOfNodes)

	dfsRecursive(headNode)
	numberOfNodesReachable++
	fmt.Println(numberOfNodes - numberOfNodesReachable)
}

func dfsRecursive(s int) {

	sIndex := s - 1
	vis[sIndex] = true

	for i := 0; i < len(graph[sIndex]); i++ {

		elementAdjacency := graph[sIndex][i]

		if elementAdjacency == 0 {
			//break
		}

		indexElementAdjacency := elementAdjacency - 1

		if !vis[indexElementAdjacency] {
			numberOfNodesReachable++
			dfsRecursive(elementAdjacency)
		}
	}
}
