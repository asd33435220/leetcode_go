package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	res := longestOnes(nums, 2)
	fmt.Println(res)
}

func longestOnes(nums []int, k int) int {
	var left, count, max, diffNumber int
	for _, value := range nums {
		if value == 0 {
			count++
		} else {
			diffNumber++
			count++
		}
		for diffNumber > k {
			if nums[left] == 1 {
				diffNumber--
			}
			count--
			left++
		}
		if count > max {
			max = count
		}
	}
	left = 0
	count = 0
	diffNumber = 0
	for _, value := range nums {
		if value == 1 {
			count++
		} else {
			diffNumber++
			count++
		}
		for diffNumber > k {
			if nums[left] == 0 {
				diffNumber--
			}
			count--
			left++
		}
		if count > max {
			max = count
		}
	}
	return max
}
