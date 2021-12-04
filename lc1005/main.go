package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{2, -3, -1, 5, -4}
	res := largestSumAfterKNegations(nums, 2)
	fmt.Println(res)
}

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		}
		res += nums[i]
		fmt.Println("nums[i]", nums[i])
	}
	if k > 0 && k%2 != 0 {
		res -= 2 * nums[len(nums)-1]
	}
	return res
}
