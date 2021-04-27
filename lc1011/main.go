package main

import "fmt"

func check(weights []int, D int, t int) bool {
	j := 0
	for i := 1; i <= D; i++ {
		total := 0
		for total <= t {
			total += weights[j]
			j++
			if j == len(weights) && total <= t {
				return true
			}
		}
		j--
	}
	return false
}
func shipWithinDays(weights []int, D int) int {
	left := weights[len(weights)-1]
	right := 0
	for i := 0; i < len(weights); i++ {
		right += weights[i]
	}
	for left < right {
		mid := (left + right) >> 1
		if check(weights, D, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func main() {
	weights := []int{1,2,3,4,5,6,7,8,9,10}
	D := 10
	res := shipWithinDays(weights, D)
	fmt.Println(res)
}
