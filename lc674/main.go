package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 5, 4, 7}
	res := findLengthOfLCIS(nums)
	fmt.Println(res)
}

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	nums = append(nums, 0)
	count := 1
	maxCount := 1
	lastNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > lastNum {
			count++
		} else {
			if count > maxCount {
				maxCount = count
			}
			count = 1
		}
		lastNum = nums[i]
	}
	return maxCount
}
