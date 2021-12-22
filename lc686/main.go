package main

import (
	"fmt"
	"strings"
)

func repeatedStringMatch(a string, b string) int {
	myMap := make(map[uint8]bool)
	for i := 0; i < len(b); i++ {
		myMap[b[i]] = true
	}
	for i := 0; i < len(a); i++ {
		myMap[a[i]] = false
	}
	for _, value := range myMap {
		if value {
			return -1
		}
	}
	times := len(b)/len(a) + 1
	str := ""
	for i := 0; i < times-1; i++ {
		str += a
	}
	for i := times - 1; i < times+2; i++ {
		if strings.Contains(str, b) {
			return i
		}
		str += a
	}
	return -1
}
func main() {
	a := "abc"
	b := "cabcabca"
	res := repeatedStringMatch(a, b)
	fmt.Println("res", res)
}
