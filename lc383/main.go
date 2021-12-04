package main

import (
	"fmt"
)

func main() {
	res := canConstruct("aa", "aab")
	fmt.Println("Res", res)
}

func canConstruct(ransomNote string, magazine string) bool {
	mMap := make(map[uint8]int)
	for i := 0; i < len(magazine); i++ {
		mMap[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		char := ransomNote[i]
		if mMap[char] == 0 {
			return false
		} else {
			mMap[char]--
		}
	}
	return true
}
