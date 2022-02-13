package main

import (
	"fmt"
)

func main() {
	res := maxNumberOfBalloons("nlaebolko")
	fmt.Println(res)
}
func maxNumberOfBalloons(text string) (ans int) {
	wordMap := make(map[byte]int)
	for i := 0; i < len(text); i++ {
		wordMap[text[i]]++
	}
	for true {
		if addCount(wordMap) {
			ans++
		} else {
			return
		}
	}
	return
}

func addCount(wordMap map[byte]int) bool {
	text := []byte("balloon")
	for _, value := range text {
		if wordMap[value] <= 0 {
			return false
		} else {
			wordMap[value]--
		}
	}
	return true
}
