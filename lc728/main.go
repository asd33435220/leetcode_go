package main

import (
	"fmt"
)

func selfDividingNumbers(left int, right int) (res []int) {
	check := func(number int) bool {
		num := number
		for num > 0 {
			curr := num % 10
			if curr == 0 || number%curr != 0 {
				return false
			}
			num = num / 10
		}
		return true
	}
	for i := left; i <= right; i++ {
		if check(i) {
			res = append(res, i)
		}
	}
	return
}
func main() {

	res := selfDividingNumbers(20, 22)
	fmt.Println("res", res)
}
