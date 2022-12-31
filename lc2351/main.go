package main

import "fmt"

func repeatedCharacter(s string) byte {
	myMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		myMap[s[i]]++
		if myMap[s[i]] == 2 {
			return s[i]
		}
	}
	return 0
}
func main() {
	fmt.Println(repeatedCharacter("abccbaacz"))
}
