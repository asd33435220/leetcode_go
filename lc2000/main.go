package main

import (
	"fmt"
)

func main() {
	word := "acefd"
	ch := 'd'
	res := reversePrefix(word, byte(ch))
	fmt.Println("Res", res)

}

func reversePrefix(word string, ch byte) string {
	index := 0
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			index = i
			break
		}
	}
	list := []byte(word)
	prefix := list[:index+1]
	prefix = reverse(prefix)
	return string(append(prefix, list[index+1:]...))

}
func reverse(str []byte) []byte {
	n := len(str)
	for i := 0; i < n/2; i++ {
		str[i], str[n-i-1] = str[n-i-1], str[i]
	}
	return str
}
