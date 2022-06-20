package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 3, 2, 6}
	res := maxRotateFunction(nums)
	fmt.Println(res)
}

func maxRotateFunction(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	f := 0
	for i, num := range nums {
		f += i * num
	}
	max := f
	for i := len(nums) - 1; i > 0; i-- {
		f += sum - len(nums)*nums[i]
		if f > max {
			max = f
		}
	}
	return max
}
