package main

import "fmt"

func main() {
	A := []int{0, 4, 2, 5, 7, 3}
	res := sortArrayByParity(A)
	fmt.Println(res)
}
func sortArrayByParity(nums []int) []int {
	left := make([]int, 0)
	right := make([]int, 0)
	for _, value := range nums {
		if value%2 == 0 {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}
	return append(left, right...)
}
