package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//This tells us if the labeling is actually a beta-labeling.
func unique(intSlice []int) bool {
	keys := make(map[int]bool)
	list := []int{}
	isUnique := true
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		} else {
			isUnique = false
		}
	}
	return isUnique
}

//This generates the vertex labels.
func generateVertexLabels(numVertices int) []int {
	rand.Seed(time.Now().Unix())
	vertexSet := make([]int, numVertices)
	vertexSet = rand.Perm(numVertices)
	return vertexSet
}

//This generates the edge set but is specific to path graphs.
func generateEdgeSet(numVertices int) [][]int {
	edgeSet := make([][]int, numVertices-1)
	for i := 0; i < len(edgeSet); i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}
	return edgeSet
}

//This generates the edge labels.
func generateEdgeLabels(vertexSet []int, edgeSet [][]int) []int {
	edgeLabels := make([]int, len(edgeSet))
	for i := 0; i < len(edgeLabels); i++ {
		edgeLabels[i] = int(math.Abs(float64(vertexSet[edgeSet[i][0]] - vertexSet[edgeSet[i][1]])))
	}
	return edgeLabels
}

//This is the main function.
func main() {
	var vertexSet []int
	var edgeSet [][]int
	var edgeLabels []int
	numVertices := 6
	for {
		vertexSet = generateVertexLabels(numVertices)
		edgeSet = generateEdgeSet(numVertices)
		edgeLabels = generateEdgeLabels(vertexSet, edgeSet)
		if unique(edgeLabels) == true {
			break
		}
	}
	fmt.Println("Vertices:", vertexSet)
	fmt.Println("Edges:", edgeSet)
	fmt.Println("Edge Labels:", edgeLabels)

}
