package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(prices)
	fmt.Println(res)
}

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i-1][1]+prices[i], dp[i-1][0])
		dp[i][1] = Max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[n-1][0]
}
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
