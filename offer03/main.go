package main

import "fmt"

func main() {
	mySlice := []int{1,2,3,4,2}
	result:=findRepeatNumber(mySlice)
	fmt.Println(result)
}
func findRepeatNumber(nums []int) int {
	myMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		_, ok := myMap[nums[i]]
		if ok {
			return nums[i]
		} else {
			myMap[nums[i]] = 1
		}
	}
	return 0
}
