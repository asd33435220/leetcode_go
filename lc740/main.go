package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 3, 3, 3, 4}
	res := deleteAndEarn(nums)
	fmt.Println(res)
}
func maxInt(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	arr := make([]int, max+1)
	for i := 0; i < len(nums); i++ {
		arr[nums[i]]++
	}
	dp := make([]int, max+1)
	dp[0] = 0
	dp[1] = arr[1]
	dp[2] = maxInt(dp[1], arr[2]*2)
	for i := 2; i <= max; i++ {
		dp[i] = maxInt(dp[i-1], dp[i-2]+arr[i]*i)
	}
	return dp[len(dp)-1]
}
