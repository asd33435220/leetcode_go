package main

import "fmt"

func constructArray(n int, k int) []int {
	// k = 1 递增序列
	// k = 2
	res := make([]int, n)
	//gap := n - k + 1
	p := 1
	q := n
	for i := 0; i < k; i++ {
		if i%2 == 0 {
			res[i] = p
			p++
		} else {
			res[i] = q
			q--
		}
	}
	for i := k; i < n; i++ {
		if k%2 != 0 {
			res[i] = p
			p++
		} else {
			res[i] = q
			q--
		}
	}

	return res
}
func main() {
	fmt.Println(constructArray(5, 3))
}
