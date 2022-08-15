package main

import "fmt"

func maxScore(s string) int {
	count := 0
	if s[0] == '0' {
		count++
	}
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			count++
		}
	}
	curr, max := count, count
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			curr++
		} else {
			curr--
		}
		if curr > max {
			max = curr
		}
	}
	return max
}
func main() {
	res := maxScore("011010")
	fmt.Println("Res", res)
}
