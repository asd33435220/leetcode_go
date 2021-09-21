package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	wordArr := strings.Split(s, " ")
	newArr := make([]string, 0)
	for _, word := range wordArr {
		if word != "" {
			newArr = append(newArr, word)
		}
	}
	return len(newArr[len(newArr)-1])
}
func main() {
	s := "   fly me   to   the moon  "
	res := lengthOfLastWord(s)
	fmt.Println("res", res)
}
