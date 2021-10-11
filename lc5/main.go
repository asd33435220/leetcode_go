package main

import "fmt"

func main() {
	s := "babad"
	res := longestPalindrome(s)
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	n := len(s)
	res := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	return res
}
