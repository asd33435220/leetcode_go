package main

import "fmt"

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		myMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		value, ok := myMap[target-nums[i]]
		if ok && value != i {
			return []int{value, i}
		}
	}
	return result
}
func main() {
	nums := []int{3,2,4}
	res := twoSum(nums, 6)
	fmt.Println("res", res)
}
