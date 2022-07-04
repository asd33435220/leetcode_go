package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{4, 2, 1, 3}
	res := minimumAbsDifference(arr)
	fmt.Println("res", res)
}
func minimumAbsDifference(arr []int) (res [][]int) {
	sort.Ints(arr)
	min := arr[1] - arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < min {
			res = [][]int{}
			min = arr[i] - arr[i-1]
			res = append(res, []int{arr[i-1], arr[i]})
		} else if arr[i]-arr[i-1] == min {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}

	return
}
