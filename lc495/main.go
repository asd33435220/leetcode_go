package main

import (
	"fmt"
)

func findPoisonedDuration(timeSeries []int, duration int) (res int) {
	for i := 1; i < len(timeSeries); i++ {
		timeGap := timeSeries[i] - timeSeries[i-1]
		if timeGap > duration {
			res += duration
		} else {
			res += timeGap
		}
	}
	res += duration
	return
}
func main() {
	timeSeries := []int{1, 2, 4}
	res := findPoisonedDuration(timeSeries, 2)
	fmt.Println("res", res)
}
