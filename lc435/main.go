package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{[]int{1, 2}, []int{1, 2}, []int{1, 2}}
	res := eraseOverlapIntervals(arr)
	fmt.Println(res)
}

func eraseOverlapIntervals(intervals [][]int) int {
	res := 1
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	right := intervals[0][1]
	for j:=1;j< len(intervals);j++ {
		if intervals[j][0]>=right {
			res++
			right=intervals[j][1]
		}
	}

	return n-res
}
