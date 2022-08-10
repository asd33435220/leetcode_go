package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Number struct {
	isX   bool
	value int
}

func getNumber(str string) (num Number) {
	if str == "+x" {
		str = "+1x"
	}
	if str == "-x" {
		str = "-1x"
	}
	bytes := []byte(str)
	num.isX = bytes[len(bytes)-1] == 'x'
	if num.isX {

		val, _ := strconv.Atoi(string(bytes[:len(bytes)-1]))
		num.value = val
	} else {
		val, _ := strconv.Atoi(string(bytes))
		num.value = val
	}
	return
}
func getNumberArr(str string) (res []Number) {
	bytes := []byte(str)
	lastIndex := 0
	for i := 1; i < len(bytes); i++ {
		if bytes[i] == '-' || bytes[i] == '+' {
			number := getNumber(string(bytes[lastIndex:i]))
			lastIndex = i
			res = append(res, number)
		}
	}
	number := getNumber(string(bytes[lastIndex:]))
	res = append(res, number)
	return
}
func solveEquation(equation string) string {
	arr := strings.Split(equation, "=")
	left := arr[0]
	right := arr[1]
	if !strings.HasPrefix(left, "-") {
		left = "+" + left
	}
	if !strings.HasPrefix(right, "-") {
		right = "+" + right
	}
	leftVec := getNumberArr(left)
	rightVec := getNumberArr(right)
	leftX, leftNum, rightX, rightNum := 0, 0, 0, 0
	for _, num := range leftVec {
		if num.isX {
			leftX += num.value
		} else {
			leftNum += num.value
		}
	}
	for _, num := range rightVec {
		if num.isX {
			rightX += num.value
		} else {
			rightNum += num.value
		}
	}
	if leftX == rightX {
		if leftNum == rightNum {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	} else {
		return fmt.Sprintf("x=%d", (leftNum-rightNum)/(rightX-leftX))
	}
}
func main() {

	res := solveEquation("x+5-3+x=6+x-2")
	fmt.Println(res)
}
