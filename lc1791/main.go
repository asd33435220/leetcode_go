package main

import (
	"fmt"
)

func main() {
	isWater := [][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}
	res := findCenter(isWater)
	fmt.Println("res", res)
}

func findCenter(edges [][]int) int {
	myMap := make(map[int]int)
	for _, edge := range edges {
		myMap[edge[0]]++
		myMap[edge[1]]++
	}
	count := len(myMap) - 1
	for key, value := range myMap {
		if value == count {
			return key
		}
	}
	return -1
}
