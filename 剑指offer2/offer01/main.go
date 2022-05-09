package main

import (
	"fmt"
	"math"
)

func divide(a int, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	sign := 1
	if a > 0 && b < 0 || a < 0 && b > 0 {
		sign = -1
	}
	a = abs(a)
	b = abs(b)
	res := 0
	sum := 0
	for sum <= a {
		sum += b
		res++
	}
	if abs(res-1) > math.MaxInt32 {
		return math.MaxInt32 * sign
	} else {
		return (res - 1) * sign
	}
}
func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
func main() {
	res := divide(-2147483648, 1)
	fmt.Println(res)
}
