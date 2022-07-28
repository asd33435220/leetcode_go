package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := fractionAddition("-1/2+1/2+1/3")
	fmt.Println("res", res)
}

type Fraction struct {
	de         int
	mo         int
	isPositive bool
}

func addFraction(x, y Fraction) Fraction {
	if x.mo == 0 {
		return y
	} else if y.mo == 0 {
		return x
	}
	result := Fraction{}
	if !x.isPositive && y.isPositive {
		x, y = y, x
	}
	if x.isPositive == y.isPositive {
		result.isPositive = x.isPositive
		result.de = x.de * y.de
		result.mo = x.mo*y.de + x.de*y.mo
	} else {
		result.isPositive = true
		result.de = x.de * y.de
		result.mo = x.mo*y.de - x.de*y.mo
		if result.mo < 0 {
			result.mo = -result.mo
			result.isPositive = false
		}
	}
	if result.mo <= 1 {
		if result.mo == 0 {
			result.de = 1
		}
		return result
	}
	mes := measure(result.mo, result.de)
	if mes != 0 {
		result.de = result.de / mes
		result.mo = result.mo / mes
	}
	return result
}
func measure(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a

}

func fractionAddition(expression string) string {
	expressionArr := []byte(expression)
	if expressionArr[0] != '-' {
		expressionArr = append([]byte{'+'}, expressionArr...)
	}
	fractions := make([]Fraction, 0)
	lastIndex := 0
	for i := 1; i < len(expressionArr); i++ {
		if expressionArr[i] == '+' || expressionArr[i] == '-' {
			fractions = append(fractions, getFraction(expressionArr[lastIndex:i]))
			lastIndex = i
		}
	}
	fractions = append(fractions, getFraction(expressionArr[lastIndex:]))
	res := Fraction{
		de:         1,
		mo:         0,
		isPositive: true,
	}
	for _, f := range fractions {
		res = addFraction(res, f)
	}
	if res.isPositive {
		return fmt.Sprintf("%s/%s", strconv.Itoa(res.mo), strconv.Itoa(res.de))
	} else {
		return fmt.Sprintf("-%s/%s", strconv.Itoa(res.mo), strconv.Itoa(res.de))
	}
}
func getFraction(numberArr []byte) Fraction {
	result := Fraction{
		de:         0,
		mo:         0,
		isPositive: numberArr[0] == '+',
	}
	arr := strings.Split(string(numberArr[1:]), "/")
	result.mo, _ = strconv.Atoi(arr[0])
	result.de, _ = strconv.Atoi(arr[1])
	return result
}
