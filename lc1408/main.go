package main

import (
	"fmt"
	"strings"
)

func stringMatching(words []string) (res []string) {
out:
	for _, word1 := range words {
		for _, word2 := range words {
			if word1 != word2 && strings.Contains(word2, word1) {
				res = append(res, word1)
				continue out
			}
		}
	}
	return
}

func main() {
	words := []string{"mass", "as", "hero", "superhero"}
	res := stringMatching(words)
	fmt.Println("Res", res)
}
