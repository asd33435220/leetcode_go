package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if dp[i-coin] == math.MaxInt32 {
				continue
			}
			dp[i] = min(dp[i-coin]+1, dp[i])
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
func main() {
	res := coinChange([]int{3, 7, 9, 11}, 40)
	fmt.Println("res", res)
}
