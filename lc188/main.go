package main

import (
	"fmt"
)

func main() {
	prices := []int{}
	res := maxProfit(2, prices)
	fmt.Println(res)
}

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	//num := 2*k + 1
	var dp [][]int
	if n/2 < k {
		k = n / 2
	}
	for i := 0; i <= 2*k; i++ {
		level := make([]int, n)
		dp = append(dp, level)
		if i%2 == 0 {
			dp[i][0] = 0
		} else {
			dp[i][0] = -prices[0]
		}
	}

	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1]
		for j := 1; j <= 2*k; j++ {
			if j%2 == 0 {
				dp[j][i] = Max(dp[j][i-1], dp[j-1][i-1]+prices[i])
			} else {
				dp[j][i] = Max(dp[j][i-1], dp[j-1][i-1]-prices[i])
			}
		}
	}
	max := 0
	for j := 2; j <= 2*k; j++ {
		max = Max(max, dp[j][n-1])
	}
	return max
}
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
