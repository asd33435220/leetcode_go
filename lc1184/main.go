package main

import (
	"fmt"
)

func main() {
	dis := []int{1, 2, 3, 4}
	res := distanceBetweenBusStops(dis, 0, 1)
	fmt.Println(res)
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	var d1, d2 int
	n := len(distance)
	for i := start; i < destination; i++ {
		d1 += distance[i]
	}

	for i := destination; i < start+n; i++ {
		d2 += distance[i%n]
	}
	return min(d1, d2)
}
