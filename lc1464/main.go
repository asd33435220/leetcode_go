package main

import "fmt"

func maxProduct(nums []int) int {
	var a, b int
	for _, v := range nums {
		if v > a {
			b = a
			a = v
		} else if v > b {
			b = v
		}
	}
	return (a - 1) * (b - 1)
}
func main() {
	target := []int{1, 2, 3, 4}
	fmt.Println("res", maxProduct(target))
}
