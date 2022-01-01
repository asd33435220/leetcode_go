package main

import (
	"fmt"
)

func main() {
	res := checkPerfectNumber(18)
	fmt.Println("res", res)
}

func checkPerfectNumber(num int) bool {
	return num == 6 || num == 28 || num == 496 || num == 8128 || num == 33550336
}
