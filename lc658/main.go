package main

import (
	"fmt"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	right := sort.SearchInts(arr, x)
	left := right - 1
	for ; k > 0; k-- {
		if left < 0 {
			right++
		} else if right >= len(arr) || x-arr[left] <= arr[right]-x {
			left--
		} else {
			right++
		}
	}
	return arr[left+1 : right]
}
func main() {

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(findClosestElements(arr, 4, 3))
}
