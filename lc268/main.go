package main

import "fmt"

func missingNumber(nums []int) int {
	resultMap := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		resultMap[nums[i]] = 1
	}
	for i := 0; i < len(resultMap); i++ {
		if resultMap[i] == 0 {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{0, 1, 3}
	res := missingNumber(nums)
	fmt.Println(res)
}
