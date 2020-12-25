package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "cbacdcbc"
	result := removeDuplicateLetters(s)
	fmt.Println(result)
}

func removeDuplicateLetters(s string) string {
	myMap := setCharMap(s)

	var stack []string
	for _, runeChar := range s {
		char := string(runeChar)
		myMap[char]--
		if indexOf(stack, char) != -1 {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > char && myMap[stack[len(stack)-1]] > 0 {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, char)
	}
	return strings.Join(stack, "")

}
func setCharMap(s string) map[string]int {
	var myMap = make(map[string]int)
	for _, runeChar := range s {
		char := string(runeChar)
		value, ok := myMap[char]
		if ok {
			myMap[char] = value + 1
		} else {
			myMap[char] = 1
		}
	}
	return myMap
}
func indexOf(array []string, target string) int {
	for index, _ := range array {
		if target == array[index] {
			return index
		}
	}
	return -1
}
