package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	res := pivotIndex(nums)
	fmt.Println(res)
}

func pivotIndex(nums []int) int {
	total := 0
	sum := 0
	for _, value := range nums {
		total += value
	}
	for key, value := range nums {
		if total-sum-value == sum {
			return key
		}
		sum += value
	}
	return -1
}
