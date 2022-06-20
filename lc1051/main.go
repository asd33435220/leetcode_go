package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) (num int) {
	expects := make([]int, len(heights))
	copy(expects, heights)
	sort.Ints(expects)
	for i := 0; i < len(heights); i++ {
		if heights[i] != expects[i] {
			num++
		}
	}
	return num
}
func main() {
	heights := []int{1, 1, 4, 2, 1, 3}
	res := heightChecker(heights)
	fmt.Print(res)
}
