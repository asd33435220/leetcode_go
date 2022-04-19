package main

import (
	"fmt"
	"math"
)

func main() {
	res := shortestToChar("loveleetcode", 'v')
	fmt.Println(res)
}

func shortestToChar(s string, c byte) []int {
	n := len(s)
	res := make([]int, n)
	lastIndex := math.MaxInt
	for i := 0; i < n; i++ {
		char := s[i]
		if char == c {
			res[i] = 0
			lastIndex = i
		} else {
			res[i] = abs(i - lastIndex)
		}
	}

	lastIndex = math.MaxInt
	for i := n - 1; i >= 0; i-- {
		char := s[i]
		if char == c {
			res[i] = 0
			lastIndex = i
		} else {
			dis := abs(i - lastIndex)
			if dis < res[i] {
				res[i] = dis
			}
		}
	}
	return res

}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x

	}
}
