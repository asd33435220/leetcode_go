package main

import "fmt"

func runningSum(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = nums[0]
	for i := 1; i < len(result); i++ {
		result[i] = nums[i] + result[i-1]
	}
	return result
}

func main()  {
	nums := []int{1,2,3,4}
	res := runningSum(nums)
	fmt.Println("res",res)
}
