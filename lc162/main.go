package main

import "math"

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1

	get := func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt64
		}
		return nums[i]
	}

	for left <= right {
		mid := left + (right-left)/2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return -1
}
func main() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	findPeakElement(nums)
}
