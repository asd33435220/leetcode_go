package main

import (
	"fmt"
)

func checkValidString(s string) bool {

	leftStack := make([]int, 0)
	starsStack := make([]int, 0)
	for index, c := range s {
		if c == '(' {
			leftStack = append(leftStack, index)
		} else if c == ')' {
			if len(leftStack) > 0 {
				leftStack = leftStack[:len(leftStack)-1]
			} else if len(starsStack) > 0 {
				starsStack = starsStack[:len(starsStack)-1]
			} else {
				return false
			}
		} else if c == '*' {
			starsStack = append(starsStack, index)
		}
	}
	for len(starsStack) > 0 && len(leftStack) > 0 {
		startIndex := starsStack[0]
		leftIndex := leftStack[0]
		if startIndex < leftIndex {
			return false
		}
		starsStack = starsStack[:len(starsStack)-1]
		leftStack = leftStack[:len(leftStack)-1]
	}
	return true
}
func main() {
	res := checkValidString("((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()")
	fmt.Println("res", res)
}
