package main

import (
	"fmt"
	"strings"
)

func main() {

	res := maxDepth("()")
	fmt.Println("res", res)
}

func maxDepth(s string) int {
	deepth := 0
	max := 0
	s = strings.ReplaceAll(s, "()", "( )")
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			deepth++
		} else if s[i] == ')' {
			deepth--
		} else {
			if deepth > max {
				max = deepth
			}
		}
	}
	return max
}
