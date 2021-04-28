package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	var b uint64 = uint64(math.Floor(math.Sqrt(float64(c))))
	var a uint64 = 0
	for a <= b {
		if (a*a + b*b) < uint64(c) {
			a++
		} else if (a*a + b*b) > uint64(c) {
			b--
		} else {
			return true
		}
	}
	return false
}
func main() {
	res := judgeSquareSum(5)
	fmt.Println(res)
}
