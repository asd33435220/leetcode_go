package main

import (
	"fmt"
)

func main() {
	res := distributeCandies([]int{1, 1, 2, 2, 3, 3})
	fmt.Println("res", res)
}
func distributeCandies(candyType []int) int {
	candyMap := make(map[int]int)
	for i := 0; i < len(candyType); i++ {
		candyMap[candyType[i]]++
	}
	typeNumber := len(candyMap)
	if typeNumber > len(candyType)/2 {
		return len(candyType) / 2
	} else {
		return typeNumber
	}
}
