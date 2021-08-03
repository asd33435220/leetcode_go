package main

import (
	"fmt"
	"sort"
)

func main() {
	res := findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15})
	fmt.Println("res", res)
}
func findUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	res := 0
	arr := make([]int, len(nums))
	copy(arr, nums)
	sort.Ints(nums)
	for i := 0; i < len(arr); i++ {
		if arr[i] == nums[i] {
			res++
		} else {
			break
		}
	}
	if res == len(arr) {
		return 0
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == nums[i] {
			res++
		} else {
			break
		}
	}
	return len(arr) - res
}
