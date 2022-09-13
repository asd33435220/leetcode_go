package main

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	str := strconv.Itoa(num)
	n := len(str)
	maxIndexArr := make([]int, n)
	max := '0'
	curr := n - 1
	for i := n - 1; i >= 0; i-- {
		c := rune(str[i])
		if c > max {
			max = c
			curr = i
		}
		maxIndexArr[i] = curr
	}
	for i := 0; i < n; i++ {
		if i != maxIndexArr[i] && str[i] != str[maxIndexArr[i]] {
			arr := []byte(str)
			arr[i], arr[maxIndexArr[i]] = arr[maxIndexArr[i]], arr[i]
			str = string(arr)
			res, _ := strconv.Atoi(str)
			return res
		}
	}
	return num
}

func main() {

	fmt.Println("Res", maximumSwap(983682))
}
