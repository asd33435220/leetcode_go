package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{20, 100, 10, 12, 5, 13}
	res := increasingTriplet(nums)
	fmt.Println(res)
}
func increasingTriplet(nums []int) bool {
	n := len(nums)
	minList := make([]int, n)
	maxList := make([]int, n)
	max := 0
	min := math.MaxInt64
	for i := 0; i < n; i++ {
		if nums[i] < min {
			min = nums[i]
		}
		minList[i] = min
	}
	for i := n - 1; i >= 0; i-- {
		if nums[i] > max {
			max = nums[i]
		}
		maxList[i] = max

	}
	for i := 1; i < n-1; i++ {
		if minList[i-1] < nums[i] && nums[i] < maxList[i+1] {
			return true
		}
	}
	return false
}
