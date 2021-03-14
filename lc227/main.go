package main

import "fmt"

func calculate(s string) int {
	res := 0
	preSign := '+'
	n := len(s)
	num := 0
	var stack []int
	var isNumber = func(char uint8) bool {
		if char-'0' >= 0 && char-'0' <= 9 {
			return true
		} else {
			return false
		}
	}
	for i := 0; i < n; i++ {
		if isNumber(s[i]) {
			num = num*10 + int(s[i]-'0')
		}
		if !isNumber(s[i]) && s[i] != ' ' || i == n-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
				break
			case '-':
				stack = append(stack, -num)

				break
			case '*':
				stack[len(stack)-1] *= num
				break
			case '/':
				stack[len(stack)-1] /= num
				break
			}
			preSign = int32(s[i])
			num = 0
		}
	}
	for _, value := range stack {
		res += value
	}
	return res
}
func main() {

	s := "- 3"
	res := calculate(s)
	fmt.Println(res)
}
