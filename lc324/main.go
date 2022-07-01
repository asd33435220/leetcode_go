package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 5, 5, 6, 7}
	wiggleSort(nums)
	fmt.Println(nums)
}
func wiggleSort(nums []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	n := len(nums)
	res := make([]int, n)

	mid := n >> 1
	bigArr := nums[:mid]
	smallArr := nums[mid:]

	for i := 0; i < n; i++ {
		j := i / 2
		if i%2 == 0 {
			res[i] = smallArr[j]
		} else {
			res[i] = bigArr[j]
		}
	}
	copy(nums, res)
}
