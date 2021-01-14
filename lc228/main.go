package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	var res []string
	if len(nums) == 0 {
		return res
	}
	nums = append(nums, nums[0])
	count := 0
	lastNumber := nums[0]
	startNumber := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == lastNumber+1 {
			count++
		} else {
			if count == 0 {
				res = append(res, strconv.Itoa(lastNumber))
				startNumber = nums[i]
			} else {
				str := strconv.Itoa(startNumber) + "->" + strconv.Itoa(lastNumber)
				res = append(res, str)
				startNumber = nums[i]
				count = 0
			}
		}
		lastNumber = nums[i]
	}
	return res
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	res := summaryRanges(nums)
	fmt.Println(res)
}
