package main

import "fmt"

func calculate(s string) int {
	sign := 1
	res := 0
	i := 0
	sli := []int{1}
	for i < len(s) {
		if s[i] == ' ' {
			i++
		} else if s[i] == '+' {
			sign = sli[len(sli)-1]
			i++
		} else if s[i] == '-' {
			sign = -sli[len(sli)-1]
			i++
		} else if s[i] == '(' {
			sli = append(sli, sign)
			i++
		} else if s[i] == ')' {
			sli = sli[:len(sli)-1]
			i++
		} else {
			number := 0
			for i < len(s) && s[i]-'0' <= 9 && s[i]-'0' >= 0 {
				number = number*10 + int(s[i]-'0')
				i++
			}
			res += number * sign
		}

	}
	return res
}
func main() {

	s := "- (3 + (4 + 5))"
	res := calculate(s)
	fmt.Println(res)
}
