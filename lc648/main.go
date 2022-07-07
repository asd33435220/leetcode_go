package main

import (
	"fmt"
	"sort"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	sort.Strings(dictionary)
	words := strings.Split(sentence, " ")
	for key, word := range words {
		for _, d := range dictionary {
			if strings.HasPrefix(word, d) {
				words[key] = d
				break
			}
		}
	}
	return strings.Join(words, " ")
}
func main() {

	dictionary := []string{"catt", "cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	res := replaceWords(dictionary, sentence)
	fmt.Println(res)
}
