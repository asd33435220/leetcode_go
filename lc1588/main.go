package main

import "fmt"

func sumOddLengthSubarrays(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - i
		left_even := (left + 1) / 2
		right_even := (right + 1) / 2
		left_odd := left / 2
		right_odd := right / 2
		res += (left_even*right_even + left_odd*right_odd) * nums[i]
	}
	return res

}

func main() {
	nums := []int{1, 2, 3, 4}
	res := sumOddLengthSubarrays(nums)
	fmt.Println("res", res)
}
