package main

import "fmt"

func nthUglyNumber(n int) int {
	uglyNumberArr := []int{1}
	var uglyNumber = 1
	var twoMultiple = 1
	var threeMultiple = 1
	var fiveMultiple = 1
	var twoIndex = 0
	var threeIndex = 0
	var fiveIndex = 0
	for i := 0; i < n-1; i++ {
		twoMultiple = uglyNumberArr[twoIndex] * 2
		threeMultiple = uglyNumberArr[threeIndex] * 3
		fiveMultiple = uglyNumberArr[fiveIndex] * 5
		uglyNumber = Min(twoMultiple, threeMultiple, fiveMultiple)
		if twoMultiple == uglyNumber {
			twoMultiple /= 2
			twoIndex++
		} else {
			twoMultiple /= 2
		}
		if threeMultiple == uglyNumber {
			threeMultiple /= 2
			threeIndex++
		} else {
			threeMultiple /= 2
		}
		if fiveMultiple == uglyNumber {
			fiveMultiple /= 2
			fiveIndex++
		} else {
			fiveMultiple /= 2
		}
		uglyNumberArr = append(uglyNumberArr, uglyNumber)
	}
	return uglyNumber
}
func Min(one, two, three int) int {
	if one <= two {
		if one <= three {
			return one
		} else {
			return three
		}
	} else {
		if two <= three {
			return two
		} else {
			return three
		}
	}
}
func main() {
	res := nthUglyNumber(13)
	fmt.Println(res)
}
