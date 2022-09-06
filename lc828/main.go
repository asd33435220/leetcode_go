package main

import (
	"fmt"
)

func main() {
	res := uniqueLetterString("ABCAC")
	fmt.Println(res)
}

func uniqueLetterString(s string) int {
	lastIndexMap := make([]int, 26)
	curIndexMap := make([]int, 26)
	for i := 0; i < 26; i++ {
		lastIndexMap[i] = -1
		curIndexMap[i] = -1
	}
	chars := []byte(s)
	res := 0
	for i := 0; i < len(chars); i++ {
		index := chars[i] - 'A'
		if curIndexMap[index] != -1 {
			res += (i - curIndexMap[index]) * (curIndexMap[index] - lastIndexMap[index])
		}
		lastIndexMap[index] = curIndexMap[index]
		curIndexMap[index] = i
	}
	for i := 0; i < 26; i++ {
		if curIndexMap[i] > -1 {
			res += (curIndexMap[i] - lastIndexMap[i]) * (len(s) - curIndexMap[i])
		}
	}
	return res
}
