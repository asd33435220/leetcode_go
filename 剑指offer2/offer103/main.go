package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
func coinChange(coins []int, amount int) int {
	pack := make([]int, amount+1)
	for i := 1; i < len(pack); i++ {
		pack[i] = math.MaxInt
	}
	for _, coin := range coins {
		for i := coin; i < len(pack); i++ {
			if pack[i-coin] == math.MaxInt {
				continue
			}
			pack[i] = min(pack[i], pack[i-coin]+1)
		}
	}
	if pack[amount] == math.MaxInt {
		return -1
	} else {
		return pack[amount]
	}
}
func main() {
	nums := []int{1, 2, 5}
	res := coinChange(nums, 11)
	fmt.Println(res)
}
