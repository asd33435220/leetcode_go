package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				res[i-1] = "FizzBuzz"
				continue
			}
			res[i-1] = "Fizz"
			continue
		}

		if i%5 == 0 {
			res[i-1] = "Buzz"
			continue
		}

		res[i-1] = strconv.Itoa(i)
	}
	return res
}
func main() {
	res := fizzBuzz(15)
	fmt.Println("res", res)
}
