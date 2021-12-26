package main

import (
	"fmt"
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	strArr := strings.Split(text, " ")
	result := make([]string, 0)
	for i := 0; i < len(strArr)-2; i++ {
		if strArr[i] == first && strArr[i+1] == second {
			result = append(result, strArr[i+2])
		}
	}
	return result
}

func main() {
	s := "alice is a good girl she is a good student"
	res := findOcurrences(s, "a", "good")
	fmt.Print(res)
}
