package main

import (
	"fmt"
	"strconv"
)

func getHint(secret string, guess string) string {
	Anums := 0
	Bnums := 0
	newSecret := ""
	newGuess := ""

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			Anums++
		} else {
			newSecret += string(secret[i])
			newGuess += string(guess[i])
		}
	}
	charMap := make(map[uint8]int)
	for i := 0; i < len(newSecret); i++ {
		charMap[newSecret[i]]++
	}

	for i := 0; i < len(newGuess); i++ {
		if charMap[newGuess[i]] > 0 {
			Bnums++
			charMap[newGuess[i]]--
		}
	}

	return strconv.Itoa(Anums) + "A" + strconv.Itoa(Bnums) + "B"
}
func main() {
	res := getHint("1123", "0111")
	fmt.Println(res)
}
