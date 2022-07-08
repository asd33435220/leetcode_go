package main

import "fmt"

func minCostToMoveChips(position []int) int {
	hash := make(map[int]int)
	for _, pos := range position {
		hash[pos%2]++
	}
	if hash[0] > hash[1] {
		return hash[1]
	} else {
		return hash[0]
	}
}

func main() {
	arr := []int{1, 5, 7, 8, 5, 3, 4, 2, 1, 2}
	res := minCostToMoveChips(arr)
	fmt.Println("res", res)
}
