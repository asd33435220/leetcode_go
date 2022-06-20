package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 1, 5, 4}
	res := countKDifference(nums, 2)
	fmt.Println("Res", res)

}

func countKDifference(nums []int, k int) (ans int) {
	myMap := make(map[int]int)
	for _, value := range nums {
		ans += myMap[value-k] + myMap[value+k]
		myMap[value]++
	}
	return
}
