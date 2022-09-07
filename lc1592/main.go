package main

import (
	"fmt"
	"strings"
)

func reorderSpaces(s string) (ans string) {
	words := strings.Fields(s)
	space := strings.Count(s, " ")
	lw := len(words) - 1
	if lw == 0 {
		return words[0] + strings.Repeat(" ", space)
	}
	return strings.Join(words, strings.Repeat(" ", space/lw)) + strings.Repeat(" ", space%lw)
}

func main() {
	res := reorderSpaces("a")
	fmt.Println("res", res)
}
