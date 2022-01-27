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

func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

//This is the main function.
func main() {

	numVertices := 6

	rand.Seed(time.Now().Unix())
	vertexSet := make([]int, numVertices)
	n := rand.Intn(numVertices)
	vertexSet[n] = numVertices - 1

	fmt.Println(n)

	//theTruth := false

	edgeSet := make([][]int, numVertices-1)

	for i := 0; i < len(edgeSet); i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}

	//Extremely Inefficient Code
	table := make([][]int, numVertices)
	table[0] = []int{1}
	table[1] = []int{0, 2}
	table[2] = []int{1, 3}
	table[3] = []int{2, 4}
	table[4] = []int{3, 5}
	table[5] = []int{4}

	forbiddenIndices := make([]int, 0)

	magicNumber := table[n][rand.Intn(len(table[n]))]

	vertexSet[magicNumber] = 0

	forbiddenIndices = append(forbiddenIndices, n)
	forbiddenIndices = append(forbiddenIndices, magicNumber)

	allowedNumbers := []int{1, 2, 3, 4}

	fmt.Println(forbiddenIndices)

	for i, _ := range vertexSet {
		if !contains(forbiddenIndices, i) {
			numB := rand.Intn(len(allowedNumbers))
			vertexSet[i] = allowedNumbers[numB]
			allowedNumbers = RemoveIndex(allowedNumbers, numB)
			forbiddenIndices = append(forbiddenIndices, i)
		}
	}

	edgeLabels := make([]int, len(edgeSet))
	for i := 0; i < len(edgeLabels); i++ {
		edgeLabels[i] = int(math.Abs(float64(vertexSet[edgeSet[i][0]] - vertexSet[edgeSet[i][1]])))
	}

	fmt.Println("Vertices:", vertexSet)
	fmt.Println("Edges:", edgeSet)
	fmt.Println("Edge Labels:", edgeLabels)

}
