package main

import (
	"fmt"
	"math"
)

func main() {
	vertexSet := make([]int, 6)
	vertexSet[0] = 0
	vertexSet[1] = 5
	vertexSet[2] = 1
	vertexSet[3] = 4
	vertexSet[4] = 2
	vertexSet[5] = 3

	fmt.Println("Vertices:", vertexSet)
	edgeSet := make([][]int, len(vertexSet)-1)

	for i := 0; i < len(edgeSet); i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}

	fmt.Println("Edges:", edgeSet)

	edgeLabels := make([]int, len(edgeSet))

	for i := 0; i < len(edgeLabels); i++ {
		edgeLabels[i] = int(math.Abs(float64(vertexSet[edgeSet[i][0]] - vertexSet[edgeSet[i][1]])))
	}

	fmt.Println("Edge Labels:", edgeLabels)

}
