package main

import "fmt"

func main() {
	res := lastRemaining(9)
	fmt.Println(res)
}

func lastRemaining(n int) int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	times := 0
	for len(arr) > 1 {
		if times%2 == 0 {
			for i := 0; i < len(arr); i++ {
				arr = append(arr[:i], arr[i+1:]...)
			}
		} else {
			for i := len(arr); i >= 1; i -= 2 {
				arr = append(arr[:i-1], arr[i:]...)
			}
		}
		times++
	}
	return arr[0]
}
