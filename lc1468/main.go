package main

import "fmt"

func xorOperation(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		res ^= start + i*2
	}
	return res
}

func main() {
	res := xorOperation(4, 3)
	fmt.Println("res", res)
}
