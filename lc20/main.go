package main

import "fmt"

func isValid(s string) bool {
	stack := make([]int, 0)
	byteSlice := []byte(s)
	for i := 0; i < len(byteSlice); i++ {
		str := string(byteSlice[i])
		switch str {
		case "(":
			stack = append(stack, 0)
			break
		case "[":
			stack = append(stack, 1)
			break
		case "{":
			stack = append(stack, 2)
			break
		case ")":
			if len(stack) > 0 && stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case "]":
			if len(stack) > 0 && stack[len(stack)-1] == 1 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case "}":
			if len(stack) > 0 && stack[len(stack)-1] == 2 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		}
	}
	return len(stack) == 0
}
func main() {
	res := isValid("{[]}")
	fmt.Println("res", res)
}
