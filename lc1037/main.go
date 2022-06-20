package main

import (
	"fmt"
)

func isBoomerang(points [][]int) bool {
	p0 := points[0]
	p1 := points[1]
	p2 := points[2]
	return (p1[1]-p2[1])*(p1[0]-p0[0]) != (p1[0]-p2[0])*(p1[1]-p0[1])
}
func main() {
	grid := [][]int{{1, 1}, {2, 2}, {3, 3}}
	res := isBoomerang(grid)
	fmt.Println(res)
}
