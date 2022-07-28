package main

import (
	"fmt"
	"sort"
)

func main() {

	res := arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12})
	fmt.Println(res)
}

func arrayRankTransform(arr []int) []int {
	n := len(arr)
	sortedArr := make([]int, n)
	copy(sortedArr, arr)
	sort.Ints(sortedArr)
	hash := make(map[int]int)
	for _, v := range sortedArr {
		if _, ok := hash[v]; !ok {
			hash[v] = len(hash) + 1
		}
	}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = hash[arr[i]]
	}
	return result
}
