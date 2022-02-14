package main

import "fmt"

func singleNonDuplicate(nums []int) (ans int) {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		next := 0
		if mid%2 == 0 {
			next = nums[mid+1]
		} else {
			next = nums[mid-1]
		}
		if nums[mid] == next {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func main() {
	nums := []int{1, 1, 3, 3, 5}
	res := singleNonDuplicate(nums)
	fmt.Println("res", res)
}
