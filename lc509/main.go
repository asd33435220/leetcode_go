package main

import "fmt"

func main() {
	res := fib(5)
	fmt.Println(res)
}

func fib(n int) int {
	var resultList = make([]int, n)
	resultList[0] = 1
	resultList[1] = 1
	var walk func(int) int
	walk = func(n int) int {
		if resultList[n-1] != 0 {
			return resultList[n-1]
		} else {
			res := walk(n-2) + walk(n-1)
			resultList[n-1] = res
			return res
		}
	}
	res := walk(n)
	return res
}
