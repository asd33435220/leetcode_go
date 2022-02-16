package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	res := combinationSum2(candidates, 8)
	fmt.Println("res", res)
}

func combinationSum2(candidates []int, target int) (res [][]int) {
	sort.Ints(candidates)
	for i := len(candidates) - 1; i >= 0; i-- {
		if candidates[i] <= target {
			candidates = candidates[:i+1]
			break
		}
	}
	var walk func(begin, target int)
	var array []int
	n := len(candidates)
	walk = func(begin, target int) {
		if target == 0 {
			res = append(res, append([]int{}, array...))
		}
		for i := begin; i < n; i++ {
			if target-candidates[i] < 0 {
				break
			}
			if i > begin && candidates[i] == candidates[i-1] {
				continue
			}
			array = append(array, candidates[i])
			walk(i+1, target-candidates[i])
			array = array[:len(array)-1]
		}

	}
	walk(0, target)
	return
}
