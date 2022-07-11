package main

import "fmt"

func main() {
	md := Constructor()
	md.BuildDict([]string{"hello", "leetcode"})
	res := md.Search("hello")
	fmt.Println(res)
	res = md.Search("hhllo")
	fmt.Println(res)
	res = md.Search("hell")
	fmt.Println(res)
	res = md.Search("leetcoded")
	fmt.Println(res)
}

type MagicDictionary []string

func Constructor() MagicDictionary {
	return nil
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	*this = dictionary
}

func checkWord(x, y string) bool {
	count := 0
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return count == 1
}

func (this *MagicDictionary) Search(searchWord string) bool {
	targetLen := len(searchWord)
	var targetWords []string
	for _, word := range *this {
		if len(word) == targetLen {
			targetWords = append(targetWords, word)
		}
	}
	if len(targetWords) == 0 {
		return false
	}
	for _, word := range targetWords {
		if checkWord(word, searchWord) {
			return true
		}
	}
	return false
}
