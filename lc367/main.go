package main

import "fmt"

func isPerfectSquare(num int) bool {
	left := 1
	right := num
	for left <= right {
		mid := (right + left) / 2
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false

}
func main() {
	res := isPerfectSquare(14)
	fmt.Println(res)
}
