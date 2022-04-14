package main

import (
	"fmt"
)

func main() {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}
	res := maximumWealth(accounts)

	fmt.Println("res", res)
}
func Sum(account []int) int {
	sum := 0
	for _, value := range account {
		sum += value
	}
	return sum
}
func maximumWealth(accounts [][]int) int {
	max := 0
	for _, account := range accounts {
		sum := Sum(account)
		if sum > max {
			max = sum
		}
	}
	return max
}
