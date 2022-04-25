package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	var walk func(int) int
	hashMap := make(map[int]int)

	hashMap[1] = 1
	hashMap[2] = 1
	walk = func(n int) int {
		if value, ok := hashMap[n]; ok {
			return value
		} else {
			hashMap[n] = (walk(n-1) + walk(n-2)) % 1000000007
			return hashMap[n]
		}
	}
	return walk(n)
}
func main() {
	res := fib(7)
	fmt.Println("res", res)
}
