package main

import (
	"fmt"
)

func minOperations(logs []string) int {
	level := 0
	for _, log := range logs {
		if log == "./" {
			continue
		} else if log == "../" {
			if level > 0 {
				level -= 1
			}
		} else {
			level += 1
		}
	}
	return level
}

func main() {
	res := minOperations([]string{"./", "../", "x/"})
	fmt.Println("res", res)
}
