package main

import "fmt"

func main() {

	res := diStringMatch("III")
	fmt.Println(res)
}

func diStringMatch(s string) (res []int) {
	n := len(s) + 1
	res = make([]int, n)
	low, high := 0, n-1
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == 'I' {
			res[i] = low
			low++
		} else {
			res[i] = high
			high--
		}
	}

	if s[n-2] == 'I' {
		res[n-1] = high
	} else {
		res[n-1] = low
	}
	return
}
