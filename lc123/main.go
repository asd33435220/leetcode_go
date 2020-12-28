package main

import "fmt"

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	res := maxProfit(prices)
	fmt.Println(res)
}

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][5]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = Max(dp[i-1][0]-prices[i], dp[i-1][1])
		dp[i][2] = Max(dp[i-1][1]+prices[i], dp[i-1][2])
		dp[i][3] = Max(dp[i-1][2]-prices[i], dp[i-1][3])
		dp[i][4] = Max(dp[i-1][3]+prices[i], dp[i-1][4])
	}
	return Max(dp[n-1][2], dp[n-1][4])
}
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
