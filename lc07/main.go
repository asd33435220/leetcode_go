package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverseArray(arr []string) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
func reverse(x int) int {
	var str string
	if x < 0 {
		str = strconv.Itoa(-x)
	} else {
		str = strconv.Itoa(x)
	}
	arr := strings.Split(str, "")
	reverseArray(arr)
	str = strings.Join(arr, "")
	maxInt := strconv.Itoa(math.MaxInt32)
	if len(str) >= len(maxInt) && str > maxInt {
		return 0
	}
	resultNumber,_ := strconv.Atoi(str)
	if x < 0 {
		return -resultNumber
	} else {
		return resultNumber
	}
}
func main() {
	res := reverse(123)
	fmt.Println("res=", res)
}
