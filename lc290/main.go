package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	hashmap1 := map[byte]string{}
	hashmap2 := map[string]byte{}
	strs := strings.Split(s, " ")
	if len(strs) != len(s) {
		return false
	}
	for i, _ := range pattern {
		if hashmap1[pattern[i]] != "" && hashmap2[strs[i]] != pattern[i] || hashmap2[strs[i]] > 0 && hashmap1[pattern[i]] != strs[i] {
			return false
		}
		hashmap1[pattern[i]] = strs[i]
		hashmap2[strs[i]] = pattern[i]
	}
	return true
}
func main() {
	res := wordPattern("abba",
		"dog cat cat dog")
	fmt.Println(res)
}
