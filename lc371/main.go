package main

import "fmt"

func getSum(a int, b int) int {
	return a + b
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}

func main() {
	res := getSum(10, 12)
	fmt.Println("res", res)
}
