package main

import (
	"fmt"
)

func main() {
	res := generateTheString(3)
	fmt.Println("res", res)
}
func generateTheString(n int) string {
	res := ""
	for i := 0; i < n-1; i++ {
		res += "x"
	}
	if n%2 == 0 {
		res += "y"
	} else {
		res += "x"
	}
	return res
}
