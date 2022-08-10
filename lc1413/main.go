package main

import "fmt"

func minStartValue(nums []int) int {
	min := 1
	sum := 0
	for _, v := range nums {
		sum += v
		curr := 1 - sum
		if curr > min {
			min = curr
		}
	}
	return min
}
func main() {
	nums := []int{3, -5, 7, -10}
	res := minStartValue(nums)
	fmt.Println("Res", res)
}
