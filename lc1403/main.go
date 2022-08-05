package main

import (
	"fmt"
	"sort"
)

func minSubsequence(nums []int) []int {
	ans := make([]int, 0)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	sort.Ints(nums)
	curr := 0
	for i := len(nums) - 1; i >= 0; i-- {
		ans = append(ans, nums[i])
		curr += nums[i]
		sum -= nums[i]
		if curr > sum {
			break
		}
	}

	return ans
}
func main() {
	res := minSubsequence([]int{3, 5, 7, 6, 7})
	fmt.Println("Res", res)
}
