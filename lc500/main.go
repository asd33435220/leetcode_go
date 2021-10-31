package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	res := findWords(words)
	fmt.Println("res", res)
}

func findWords(words []string) (res []string) {
	keyMap := make(map[string]int)

	line := "qwertyuiop"
	for _, str := range strings.Split(line, "") {
		keyMap[str] = 1
	}
	line = "asdfghjkl"
	for _, str := range strings.Split(line, "") {
		keyMap[str] = 2
	}
	line = "zxcvbnm"
	for _, str := range strings.Split(line, "") {
		keyMap[str] = 3
	}

	for _, str := range words {
		word := strings.Split(strings.ToLower(str), "")
		lastNum := 0
		for key, c := range word {
			if key == 0 {
				lastNum = keyMap[c]
			}
			num := keyMap[c]
			if key == len(word)-1 && num == lastNum {
				res = append(res, str)
			}
			if lastNum != num {
				break
			}
			lastNum = num
		}
	}
	return res
}
