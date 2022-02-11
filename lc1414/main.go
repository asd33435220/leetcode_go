package main

import "fmt"

func findMinFibonacciNumbers(k int) int {
	if k <= 3 {
		return 1
	}
	list := []int{1, 1}
	curr := 2
	for true {
		n := len(list)
		curr = list[n-1] + list[n-2]
		if curr > k {
			break
		}
		list = append(list, curr)
	}
	ans := 0
	for i := len(list) - 1; k > 0; i-- {
		if k >= list[i] {
			k -= list[i]
			ans++
		}
	}
	return ans
}

func main() {
	res := findMinFibonacciNumbers(19)
	fmt.Println("Res", res)
}
