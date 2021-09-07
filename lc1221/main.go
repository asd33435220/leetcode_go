package main

import "fmt"

func balancedStringSplit(s string) int {
	result := 0
	balance := 0
	for _, char := range s {
		if char == 'L' {
			balance++
		} else {
			balance--
		}
		if balance == 0 {
			result++
		}
	}
	return result
}
func main() {
	s := "RLRRLLRLRL"
	res := balancedStringSplit(s)
	fmt.Println("res", res)
}
