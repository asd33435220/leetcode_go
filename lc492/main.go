package main

import (
	"fmt"
	"math"
)

func constructRectangle(area int) []int {
	width := int(math.Floor(math.Sqrt(float64(area)))) + 1
	rest := 1
	for rest != 0 {
		width--
		rest = area % width
	}
	return []int{area / width, width}
}
func main() {
	res := constructRectangle(40)
	fmt.Println("res", res)
}
