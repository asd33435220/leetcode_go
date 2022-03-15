package main

import (
	"fmt"
)

func addDigits(num int) int {
	for num >= 10 {
		rest := num % 10
		num = num/10 + rest
	}
	return num
}
func main() {

	res := addDigits(1234)
	fmt.Println("res", res)
}
