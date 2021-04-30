package main

import "fmt"

func singleNumber(nums []int) int {
	myMap := make(map[int]int)
	for _, value := range nums {
		myMap[value]++
	}
	for key, value := range myMap {
		if value == 1 {
			return key
		}
	}
	return -1
}
func main() {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	res := singleNumber(nums)
	fmt.Println(res)

}
