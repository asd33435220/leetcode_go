package main

import (
	"fmt"
	"strings"
)

func main() {

	res := countValidWords("pencil-sharpener. he bought 2 pencils, 3 erasers, and 1  ")
	fmt.Println("res", res)
}

func countValidWords(sentence string) (ans int) {
	strList := strings.Split(sentence, " ")
	for i := 0; i < len(strList); i++ {
		if isValid(strList[i]) {
			ans++
		}
	}
	return
}

func isValid(str string) bool {
	if str == "" {
		return false
	}
	splitIndex := -1
	signIndex := -1
	for index, value := range str {
		if value == '-' {
			if splitIndex != -1 {
				return false
			} else {
				splitIndex = index
			}
		} else if value < 47 {
			if signIndex != -1 {
				return false
			} else {
				signIndex = index
			}
		} else if value < 58 {
			return false
		} else {
			continue
		}
	}

	if signIndex != -1 && signIndex != len(str)-1 {
		return false
	}
	if splitIndex != -1 && (splitIndex == 0 || splitIndex == len(str)-1 || str[splitIndex-1] < 58 || str[splitIndex+1] < 58) {
		return false
	}
	return true
}
