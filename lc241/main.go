package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(expression string) []int {
	var walk func(str string) []int
	walk = func(str string) []int {
		if isPureNumber(expression) {
			val, _ := strconv.Atoi(expression)
			return []int{val}
		}
		res := make([]int, 0)
		for index, c := range expression {
			if c == '-' || c == '+' || c == '*' {
				lefts := diffWaysToCompute(expression[:index])
				rights := diffWaysToCompute(expression[index+1:])
				for _, left := range lefts {
					for _, right := range rights {
						addNumber := 0
						if c == '-' {
							addNumber = left - right
						} else if c == '+' {
							addNumber = left + right
						} else {
							addNumber = left * right
						}
						res = append(res, addNumber)
					}
				}
			}
		}
		return res
	}
	res := walk(expression)
	return res
}
func isPureNumber(number string) bool {
	_, err := strconv.Atoi(number)
	if err != nil {
		return false
	} else {
		return true
	}
}
func main() {

	res := diffWaysToCompute("2*3-4*5")
	fmt.Println("res", res)
}
