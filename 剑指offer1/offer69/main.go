package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	res := 0
	left, right := 1, n-2
	for left <= right {
		mid := (right + left) / 2
		if arr[mid] > arr[mid+1] {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
func main() {
	nums := []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	res := peakIndexInMountainArray(nums)
	fmt.Println("res", res)
}
