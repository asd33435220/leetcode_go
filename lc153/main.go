package main

import "fmt"

func findMin(nums []int) int {
	low := 0
	height := len(nums) - 1
	for low < height {
		pivot := low + (height-low)/2
		if nums[pivot] < nums[height] {
			height = pivot
		} else {
			low = pivot + 1
		}
	}
	return nums[low]
}
func main() {
	nums := []int{2, 3, 4, 5, 6, 8, 9, 1}
	res := findMin(nums)
	fmt.Println(res)
}
