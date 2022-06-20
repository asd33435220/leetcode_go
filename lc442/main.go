package main

import "fmt"

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for _, value := range nums {
		index := abs(value) - 1
		if nums[index] < 0 {
			res = append(res, abs(value))
		} else {
			nums[index] = -nums[index]
		}
	}
	fmt.Println("nums", nums)
	return res
}
func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
func main() {
	//[10,2,5,10,9,1,1,4,3,7]
	res := findDuplicates([]int{10, 2, 5, 10, 9, 1, 1, 4, 3, 7})
	fmt.Println("res", res)
}
