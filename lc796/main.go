package main

import (
	"fmt"
)

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i:]+s[:i] == goal {
			return true
		}
	}
	return false
}
func main() {

	res := rotateString("abcde", "cdeab")
	fmt.Println("res", res)
}
