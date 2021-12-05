package main

import (
	"fmt"
	"strings"
)

func truncateSentence(s string, k int) string {
	arr := strings.Split(s, " ")
	arr = arr[:k]
	return strings.Join(arr, " ")
}
func main() {

	res := truncateSentence("Hello how are you Contestant", 4)
	fmt.Println("res=", res)
}
