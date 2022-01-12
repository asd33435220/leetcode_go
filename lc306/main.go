package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "12122436"
	result := isAdditiveNumber(s)
	fmt.Println("max Int")
	fmt.Println(result)
}

func isAdditiveNumber(num string) bool {
	for i := 1; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if isValid(num, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func isValid(num string, i, j, k int) bool {

	num1 := string([]byte(num)[i:j])
	num2 := string([]byte(num)[j:k])
	x1, _ := strconv.Atoi(num1)
	x2, _ := strconv.Atoi(num2)
	if x1 > 0 && strings.HasPrefix(num1, "0") || x2 > 0 && strings.HasPrefix(num2, "0") {
		return false
	}
	sum := add(num1, num2)
	if len(sum) > len(num[k:]) {
		return false
	}
	for s := 0; s < len(sum); s++ {
		if sum[s] != num[s+k] {
			return false
		}
	}
	if len(sum)+k == len(num) {
		return true
	} else {
		return isValid(num, j, k, len(sum)+k)
	}
}

func add(x, y string) string {
	X, _ := strconv.Atoi(x)
	Y, _ := strconv.Atoi(y)
	return strconv.Itoa(X + Y)
}
