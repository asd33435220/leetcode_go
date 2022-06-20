package main

import (
	"fmt"
)

func main() {

	res := countNumbersWithUniqueDigits(3)
	fmt.Println(res)
}
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 0
	}
	res := 9
	for i := 9; i > 10-n; i-- {
		res *= i
	}
	return res + countNumbersWithUniqueDigits(n-1)
}
