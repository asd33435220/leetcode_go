package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return maxInt(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = maxInt(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = maxInt(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[len(dp)-1]
}
func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	nums := []int{1, 2, 3, 1}
	res := rob(nums)
	fmt.Println(res)
}
