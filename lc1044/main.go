package main

import (
	"fmt"
	"strings"
)

func main() {
	res := longestDupSubstring("banana")
	fmt.Println(res)
}
func longestDupSubstring(s string) (ans string) {
	for i := 0; i < len(s); i++ {
		end := 0
		if i+len(ans)+1 > len(s) {
			end = len(s)
		} else {
			end = i + len(ans) + 1
		}
		for strings.Contains(s[i+1:], s[i:end]) && end-i > len(ans) {
			ans = s[i:end]
			end = i + len(ans) + 1
		}
	}
	return ans
}
