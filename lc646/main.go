package main

import (
	"fmt"
	"sort"
)

func findLongestChain(pairs [][]int) (res int) {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	curr := -99999
	for _, pair := range pairs {
		if pair[0] > curr {
			res++
			curr = pair[1]
		}
	}
	return
}
func main() {
	pairs := [][]int{{1, 2}, {2, 3}, {3, 4}}
	res := findLongestChain(pairs)
	fmt.Println(res)
}
