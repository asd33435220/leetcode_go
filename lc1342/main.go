package main

import (
	"fmt"
)

func main() {
	res := numberOfSteps(3)
	fmt.Println("res", res)
}
func numberOfSteps(num int) int {
	ans := 0
	for num > 0 {
		ans++
		if num%2 == 0 {
			num = num / 2
		} else {
			num -= 1
		}
	}
	return ans
}
