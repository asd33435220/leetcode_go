package main

import (
	"fmt"
)

func main() {
	s := "abb"
	res := permuteUnique(s)
	fmt.Println(res)
}

func permuteUnique(s string) (res []string) {
	n := len(s)
	var walk func(byteArr []byte, first int)
	walk = func(byteArr []byte, first int) {
		if n == first {
			res = append(res, string(byteArr))
		}
		myMap := make(map[byte]bool)
		for i := first; i < n; i++ {
			if value := myMap[byteArr[i]]; value {
				continue
			}
			myMap[byteArr[i]] = true
			byteArr[i], byteArr[first] = byteArr[first], byteArr[i]
			walk(byteArr, first+1)
			byteArr[i], byteArr[first] = byteArr[first], byteArr[i]
		}
	}
	walk([]byte(s), 0)
	return res
}
