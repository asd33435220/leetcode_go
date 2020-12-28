package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	result := make([]int, len(prices))
	for i, value := range prices {
		if i == 0 {
			continue
		}
		if value < minPrice {
			minPrice = value
		}
		result[i] = Max(result[i-1], value-minPrice)
	}
	return result[len(result)-1]
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(prices)
	fmt.Println(res)
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
