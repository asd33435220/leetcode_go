package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	myMap := make(map[int]int)
	for key, value := range nums {
		if lastKey, ok := myMap[value]; ok && key-lastKey <= k {
			return true
		}
		myMap[value] = key
	}
	return false
}
func main() {
	nums := []int{1, 2, 3, 1}

	res := containsNearbyDuplicate(nums, 3)
	fmt.Println(res)
}
