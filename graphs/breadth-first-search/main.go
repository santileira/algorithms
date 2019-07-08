package main

import (
	"fmt"
)

func main() {

	// number of nodes in this tree
	var numberOfNodes int

	fmt.Scanf("%d", &numberOfNodes)

	// adjacency matrix
	tree := make([][]int, numberOfNodes, numberOfNodes)

	// generates the adjacency matrix
	for i := 1; i < numberOfNodes; i++ {
		var nodeFrom, nodeTo int
		fmt.Scanf("%d", &nodeFrom)
		fmt.Scanf("%d", &nodeTo)
		tree[nodeFrom-1] = append(tree[nodeFrom-1], nodeTo)
		tree[nodeTo-1] = append(tree[nodeTo-1], nodeFrom)
	}

	// You need to find out the number of nodes lying on level x.
	var numberOfLevel int
	fmt.Scanf("%d", &numberOfLevel)

	var numberOfNodesInLevel int

	nodeRoot := 1
	levelNodeRoot := 1

	// slice to mark if the node is visited
	vis := make([]bool, numberOfNodes, numberOfNodes)
	vis[0] = true

	queue := make([]int, 0, numberOfNodes)
	queue = append(queue, nodeRoot)

	// slice to determine the level of each node
	level := make([]int, numberOfNodes, numberOfNodes)
	level[0] = levelNodeRoot

	// while queue is not empty
	for len(queue) != 0 {

		// get element and update the queue
		element := queue[0]
		queue = queue[1:]

		indexElementInTree := element-1
		elementInTree := tree[indexElementInTree]

		// when pass the level break.
		if numberOfLevel < level[indexElementInTree] {
			break
		}

		// go through all the adjacent nodes
		for i := 0; i < len(elementInTree); i++ {

			// get adjacency node
			adjacencyNode := elementInTree[i]

			// when value is zero, don't exist adjacency node and the rest of elements are empty.
			if adjacencyNode == 0 {
				break
			}

			indexAdjacencyNode := adjacencyNode - 1

			// if the node is not visited
			if !vis[indexAdjacencyNode] {

				level[indexAdjacencyNode] = level[indexElementInTree] + 1
				queue = append(queue, adjacencyNode)
				vis[indexAdjacencyNode] = true

				if numberOfLevel == level[indexAdjacencyNode] {
					numberOfNodesInLevel++
				}

			}
		}
	}

	fmt.Println(numberOfNodesInLevel)

}
