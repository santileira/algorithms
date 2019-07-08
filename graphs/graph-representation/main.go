package main

import (
	"fmt"
)

func main() {

	var nodes, edges int

	fmt.Scanf("%d", &nodes)
	fmt.Scanf("%d", &edges)
	//fmt.Printf("el valor de nodes es %d", nodes)
	//fmt.Printf("el valor de edges es %d", edges )
	array := make([][]bool, nodes, nodes)

	for i := range array {
		array[i] = make([]bool, nodes) /* again the type? */
	}

	for i := 0; i < edges; i++ {
		var nodeFrom, nodeTo int
		fmt.Scanf("%d", &nodeFrom)
		fmt.Scanf("%d", &nodeTo)
		array[nodeFrom-1][nodeTo-1] = true
		array[nodeTo-1][nodeFrom-1] = true

	}

	var queries int
	fmt.Scanf("%d", &queries)

	for i := 0; i < queries; i++ {
		var nodeFrom, nodeTo int
		fmt.Scanf("%d", &nodeFrom)
		fmt.Scanf("%d", &nodeTo)
		if array[nodeFrom-1][nodeTo-1] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		//fmt.Println(array[nodeFrom-1][nodeTo-1])
	}
	//fmt.Println(array)
}
