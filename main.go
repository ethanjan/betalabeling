package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//This gives the index of an element in a slice.
func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

//This creates a new slice containing a specific range of numbers.
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

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

//This informs us if an element is contained in a slice.
func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

//This removes an element at a specific index in an array.
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

//This is the main function.
func main() {

	numVertices := 37

	rand.Seed(time.Now().Unix())
	vertexSet := make([]int, numVertices)

	edgeSet := make([][]int, numVertices-1)

	//This will generate the edges for any path graph.
	/*for i := 0; i < len(edgeSet); i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}*/

	//This will hopefully generate the edges for a binary tree with 7 vertices.
	/*edgeSet[0] = make([]int, 2)
	edgeSet[0][0] = 0
	edgeSet[0][1] = 1
	edgeSet[1] = make([]int, 2)
	edgeSet[1][0] = 0
	edgeSet[1][1] = 2
	edgeSet[2] = make([]int, 2)
	edgeSet[2][0] = 1
	edgeSet[2][1] = 3
	edgeSet[3] = make([]int, 2)
	edgeSet[3][0] = 1
	edgeSet[3][1] = 4
	edgeSet[4] = make([]int, 2)
	edgeSet[4][0] = 2
	edgeSet[4][1] = 5
	edgeSet[5] = make([]int, 2)
	edgeSet[5][0] = 2
	edgeSet[5][1] = 6*/

	/*startingIncrement := 0
	moreIncrement := 0
	for i := 0; i < len(edgeSet); i += 2 {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = startingIncrement
		edgeSet[i][1] = startingIncrement + 1 + moreIncrement
		edgeSet[i+1] = make([]int, 2)
		edgeSet[i+1][0] = startingIncrement
		edgeSet[i+1][1] = startingIncrement + 2 + moreIncrement
		startingIncrement++
		moreIncrement++
	}*/

	//This function generates the edges for a very specific type of caterpillar graph.
	/*edgeSet[0] = make([]int, 2)
	edgeSet[0][0] = 0
	edgeSet[0][1] = 1
	for i := 1; i < len(edgeSet); {
		for j := 0; j < 4; j++ {
			edgeSet[i+j] = make([]int, 2)
			edgeSet[i+j][0] = i
			edgeSet[i+j][1] = i + j + 1
		}
		i += 4
	}*/
	//This function generates the edges for an extended star graph.
	/*for i := 0; i < 5; i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}
	edgeSet[5] = make([]int, 2)
	edgeSet[5][0] = 5
	edgeSet[5][1] = 6

	edgeSet[6] = make([]int, 2)
	edgeSet[6][0] = 5
	edgeSet[6][1] = 7

	edgeSet[7] = make([]int, 2)
	edgeSet[7][0] = 7
	edgeSet[7][1] = 8

	edgeSet[8] = make([]int, 2)
	edgeSet[8][0] = 5
	edgeSet[8][1] = 9

	for i := 9; i < 11; i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}

	edgeSet[11] = make([]int, 2)
	edgeSet[11][0] = 5
	edgeSet[11][1] = 12

	for i := 12; i < len(edgeSet); i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}*/

	/*for i := 0; i < 7; i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}
	edgeSet[7] = make([]int, 2)
	edgeSet[7][0] = 7
	edgeSet[7][1] = 8

	edgeSet[8] = make([]int, 2)
	edgeSet[8][0] = 7
	edgeSet[8][1] = 9

	edgeSet[9] = make([]int, 2)
	edgeSet[9][0] = 9
	edgeSet[9][1] = 10

	edgeSet[10] = make([]int, 2)
	edgeSet[10][0] = 7
	edgeSet[10][1] = 11

	for i := 11; i < 13; i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}

	edgeSet[13] = make([]int, 2)
	edgeSet[13][0] = 7
	edgeSet[13][1] = 14

	for i := 14; i < 17; i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}

	edgeSet[17] = make([]int, 2)
	edgeSet[17][0] = 7
	edgeSet[17][1] = 18

	for i := 18; i < 22; i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}

	edgeSet[22] = make([]int, 2)
	edgeSet[22][0] = 7
	edgeSet[22][1] = 23

	for i := 23; i < len(edgeSet); i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}*/

	for i := 0; i < 8; i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}
	edgeSet[8] = make([]int, 2)
	edgeSet[8][0] = 8
	edgeSet[8][1] = 9

	edgeSet[9] = make([]int, 2)
	edgeSet[9][0] = 8
	edgeSet[9][1] = 10

	edgeSet[10] = make([]int, 2)
	edgeSet[10][0] = 10
	edgeSet[10][1] = 11

	edgeSet[11] = make([]int, 2)
	edgeSet[11][0] = 8
	edgeSet[11][1] = 12

	for i := 12; i < 14; i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}

	edgeSet[14] = make([]int, 2)
	edgeSet[14][0] = 8
	edgeSet[14][1] = 15

	for i := 15; i < 18; i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}

	edgeSet[18] = make([]int, 2)
	edgeSet[18][0] = 8
	edgeSet[18][1] = 19

	for i := 19; i < 23; i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}

	edgeSet[23] = make([]int, 2)
	edgeSet[23][0] = 8
	edgeSet[23][1] = 24

	for i := 24; i < 29; i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}

	edgeSet[29] = make([]int, 2)
	edgeSet[29][0] = 8
	edgeSet[29][1] = 30

	for i := 30; i < len(edgeSet); i++ {
		edgeSet[i] = make([]int, 2)
		edgeSet[i][0] = i
		edgeSet[i][1] = i + 1
	}

	//fmt.Println(edgeSet)

	//Extremely Inefficient Code
	/*table := make([][]int, numVertices)
	table[0] = []int{1}
	table[1] = []int{0, 2}
	table[2] = []int{1, 3}
	table[3] = []int{2, 4}
	table[4] = []int{3, 5}
	table[5] = []int{4}*/

	//fmt.Println(table)

	for {
		table := make([][]int, numVertices)

		for i, _ := range table {
			for j, _ := range edgeSet {
				if contains(edgeSet[j], i) {
					num := 0
					for _, k := range edgeSet[j] {
						if !(i == k) {
							num = k
						}
					}
					table[i] = append(table[i], num)
				}
			}
		}

		//fmt.Println(table)

		edgeLabels := make([]int, numVertices-1)
		maxDistance := numVertices - 1

		allowedNumbers := makeRange(0, maxDistance)

		forbiddenIndices := make([]int, 0)

		firstIndex := rand.Intn(numVertices)

		vertexSet[firstIndex] = maxDistance

		allowedNumbers = RemoveIndex(allowedNumbers, indexOf(vertexSet[firstIndex], allowedNumbers))

		forbiddenIndices = append(forbiddenIndices, firstIndex)

		//fmt.Println(allowedNumbers)

		secondIndex := table[firstIndex][rand.Intn(len(table[firstIndex]))]

		vertexSet[secondIndex] = int(math.Abs(float64(vertexSet[firstIndex] - maxDistance)))

		allowedNumbers = RemoveIndex(allowedNumbers, indexOf(vertexSet[secondIndex], allowedNumbers))

		forbiddenIndices = append(forbiddenIndices, secondIndex)

		//fmt.Println(allowedNumbers)

		maxDistance--
		epicIndex := -1
		//count := 0

		theLength := len(allowedNumbers)
		for len(allowedNumbers) > 0 {
			for i := 0; i < len(forbiddenIndices); i++ {
				randomNumber := forbiddenIndices[i]

				stuff := make([]int, len(table[randomNumber]))

				copy(stuff, table[randomNumber])

				coolSlice := make([]int, 0)
				for _, v := range stuff {
					if !contains(forbiddenIndices, v) {
						coolSlice = append(coolSlice, v)
					}
				}
				//fmt.Println("Forbidden Indices", forbiddenIndices)
				//fmt.Println("Cool Slice", coolSlice)
				if len(coolSlice) > 0 {
					epicIndex = coolSlice[rand.Intn(len(coolSlice))]
					/*fmt.Println("Epic Index", epicIndex)
					fmt.Println("Allowed Numbers", allowedNumbers)
					fmt.Println(maxDistance)*/
					if contains(allowedNumbers, int(math.Abs(float64(vertexSet[forbiddenIndices[i]]+maxDistance)))) {
						vertexSet[epicIndex] = int(math.Abs(float64(vertexSet[forbiddenIndices[i]] + maxDistance)))
						allowedNumbers = RemoveIndex(allowedNumbers, indexOf(vertexSet[epicIndex], allowedNumbers))
						forbiddenIndices = append(forbiddenIndices, epicIndex)
						maxDistance--
						//fmt.Println("Max Distance", maxDistance)
						break
					}
					if contains(allowedNumbers, int(math.Abs(float64(vertexSet[forbiddenIndices[i]]-maxDistance)))) {
						vertexSet[epicIndex] = int(math.Abs(float64(vertexSet[forbiddenIndices[i]] - maxDistance)))
						allowedNumbers = RemoveIndex(allowedNumbers, indexOf(vertexSet[epicIndex], allowedNumbers))
						forbiddenIndices = append(forbiddenIndices, epicIndex)
						maxDistance--
						//fmt.Println("Max Distance", maxDistance)
						break
					}
				}
			}
			/*fmt.Println("Vertices:", vertexSet)
			fmt.Println("Edges:", edgeSet)
			fmt.Println("Edge Labels:", edgeLabels)*/

			theLength--
			if theLength < -1 {
				break
			}

			/*count++
			if count > 50 {
				break
			}*/

		}

		//fmt.Println(epicIndex)

		for i := 0; i < len(vertexSet); i++ {
			if !contains(forbiddenIndices, i) {
				n := allowedNumbers[rand.Intn(len(allowedNumbers))]
				vertexSet[i] = n
				allowedNumbers = RemoveIndex(allowedNumbers, indexOf(n, allowedNumbers))
			}
		}

		/*
			for i, _ := range vertexSet {
				if !contains(forbiddenIndices, i) {
					numB := rand.Intn(len(allowedNumbers))
					vertexSet[i] = allowedNumbers[numB]
					allowedNumbers = RemoveIndex(allowedNumbers, numB)
					forbiddenIndices = append(forbiddenIndices, i)
				}
			}*/

		for i := 0; i < len(edgeLabels); i++ {
			edgeLabels[i] = int(math.Abs(float64(vertexSet[edgeSet[i][0]] - vertexSet[edgeSet[i][1]])))
		}
			
		/*fmt.Println("Vertices:", vertexSet)
		fmt.Println("Edges:", edgeSet)
		fmt.Println("Edge Labels:", edgeLabels)*/
		if unique(edgeLabels) {
			fmt.Println("Vertices:", vertexSet)
			fmt.Println("Edges:", edgeSet)
			fmt.Println("Edge Labels:", edgeLabels)
			break
		}
	}

}
