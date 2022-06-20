package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	res := sumOfUnique(nums)
	fmt.Println("res", res)
}

func sumOfUnique(nums []int) (ans int) {
	myMap := map[int]int{}
	for _, value := range nums {
		myMap[value]++
	}
	for key, value := range myMap {
		if value == 1 {
			ans += key
		}
	}
	return
}
