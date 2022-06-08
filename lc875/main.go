package main

import "fmt"

func getHour(piles []int, speed int) (hour int) {
	for _, number := range piles {
		hour += (number + speed - 1) / speed
	}
	return
}
func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, number := range piles {
		if number > max {
			max = number
		}
	}

	left := 1
	right := max

	for left < right {
		mid := (right + left) / 2
		hour := getHour(piles, mid)
		if hour > h {
			left = mid + 1

		} else {
			right = mid
		}
	}
	return left

}
func main() {
	piles := []int{3, 6, 7, 11}
	res := minEatingSpeed(piles, 8)
	fmt.Println("res", res)
}
