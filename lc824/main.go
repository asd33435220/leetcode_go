package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "I speak Goat Latin"
	res := toGoatLatin(sentence)
	fmt.Println(res)
}

func toGoatLatin(sentence string) string {
	wordArr := strings.Split(sentence, " ")
	res := make([]string, 0)
	for key, word := range wordArr {
		bytes := []byte(word)
		beginByte := bytes[0]
		if index := strings.IndexByte("aeiouAEIOU", beginByte); index != -1 {
			bytes = append(bytes, 'm', 'a')

		} else {
			bytes = bytes[1:]
			bytes = append(bytes, beginByte, 'm', 'a')

		}
		for i := 0; i < key+1; i++ {
			bytes = append(bytes, 'a')
		}
		res = append(res, string(bytes))
	}
	return strings.Join(res, " ")
}
