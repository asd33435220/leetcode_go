package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{9, 7}
	res := minimumDifference(nums, 2)
	fmt.Println("Res", res)

}

func minimumDifference(nums []int, k int) int {
	if len(nums) <= k {
		min := math.MaxInt64
		max := -1
		for _, value := range nums {
			if value < min {
				min = value
			}
			if value > max {
				max = value
			}
		}
		return max - min
	}

	sort.Ints(nums)
	minGap := nums[k-1] - nums[0]
	for i := 1; k+i-1 < len(nums); i++ {
		newGap := nums[k+i-1] - nums[i]
		if newGap < minGap {
			minGap = newGap
		}
	}
	return minGap
}
