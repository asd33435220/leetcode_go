package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abb"
	res := permuteUnique(s)
	fmt.Println(res)
}

func permuteUnique(s string) []string {
	var result []string
	var walk func(s string, first int, length int)
	walk = func(s string, first int, length int) {
		if first == length {
			result = append(result, s)
		}
		myMap := make(map[byte]bool)
		for i := first; i < length; i++ {
			if res := myMap[s[i]]; res {
				continue
			}
			myMap[s[i]] = true
			Arr := make([]string, 0)
			for j := 0; j < len(s); j++ {
				Arr = append(Arr, string(s[j]))
			}
			Arr[first], Arr[i] = Arr[i], Arr[first]
			S := strings.Join(Arr, "")
			walk(S, first+1, length)
		}

	}
	walk(s, 0, len(s))
	return result
}
