package main

import (
	"container/heap"
	"fmt"
)

var (
	marked []bool
	adj    [][]*point
)

// point represents vertexTo and weight to this vertex.
type point struct {
	vertexTo int
	weight   int
}

// An Element is something we manage in a priority queue.
type Element struct {
	value    int // The value of the Element; arbitrary.
	priority int // The priority of the Element in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the Element in the heap.
}

// priorityQueue holds
var priorityQueue PriorityQueue

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Element

// Len returns length of the queue
func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use less than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Element)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Element in the queue.
func (pq *PriorityQueue) update(item *Element, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	// number of vertices and edges in this algorithms
	var numberOfVertices, numberOfEdges int

	fmt.Scanf("%d", &numberOfVertices)
	fmt.Scanf("%d", &numberOfEdges)

	// initializes priority queue
	priorityQueue = make(PriorityQueue, 0)
	heap.Init(&priorityQueue)

	marked = make([]bool, numberOfVertices, numberOfVertices)
	adj = make([][]*point, numberOfVertices, numberOfVertices)

	var vertexFrom, vertexTo, weight int
	for i := 0; i < numberOfEdges; i++ {

		fmt.Scanf("%d", &vertexFrom)
		fmt.Scanf("%d", &vertexTo)
		fmt.Scanf("%d", &weight)

		pointVertexTo := &point{
			vertexTo: vertexTo,
			weight:   weight,
		}

		pointVertexFrom := &point{
			vertexTo: vertexFrom,
			weight:   weight,
		}

		adj[vertexFrom-1] = append(adj[vertexFrom-1], pointVertexTo)
		adj[vertexTo-1] = append(adj[vertexTo-1], pointVertexFrom)

	}

	minimumCost := prim(vertexFrom)

	fmt.Println(minimumCost)
}

func prim(x int) (minimumCost int) {

	element := &Element{
		value:    x,
		priority: 0,
	}
	heap.Push(&priorityQueue, element)

	for priorityQueue.Len() != 0 {
		element = heap.Pop(&priorityQueue).(*Element)
		elementValue := element.value
		indexElementValue := elementValue - 1
		if marked[indexElementValue] {
			continue
		}
		marked[indexElementValue] = true

		minimumCost += element.priority

		for i := 0; i < len(adj[indexElementValue]); i++ {
			elementAdjacency := adj[indexElementValue][i]

			if elementAdjacency == nil {
				break
			}

			newElement := &Element{
				value:    elementAdjacency.vertexTo,
				priority: elementAdjacency.weight,
			}

			if !marked[newElement.value-1] {
				heap.Push(&priorityQueue, newElement)
			}
		}

	}

	return
}
