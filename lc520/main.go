package main

import "fmt"

func detectCapitalUse(word string) bool {
	isFirstBig := word[0] < 96
	if len(word) >= 2 && !isFirstBig && word[1] < 96 {
		return false
	}

	for i := 2; i < len(word); i++ {
		if (word[i] < 96 && word[1] > 96) ||
			(word[i] > 96 && word[1] < 96) {
			return false
		}
	}
	return true
}

func main() {

	res := detectCapitalUse("Test")
	fmt.Println("res", res)
}
