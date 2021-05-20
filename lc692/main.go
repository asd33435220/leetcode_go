package main

import (
	"fmt"
	"sort"
)

type wordWithTimes struct {
	word  string
	times int
}

func topKFrequent(words []string, k int) []string {
	wordList := make([]wordWithTimes, 0)
	myMap := make(map[string]int)
	resultList := make([]string,k)
	for _, value := range words {
		myMap[value]++
	}
	for key, value := range myMap {
		wordList = append(wordList, wordWithTimes{
			word:  key,
			times: value,
		})
	}
	sort.Slice(wordList, func(i, j int) bool {
		if wordList[i].times != wordList[j].times {
			return wordList[i].times > wordList[j].times
		}else{
			return wordList[i].word < wordList[j].word
		}
	})
	for i := 0; i < k; i++ {
		resultList[i] = wordList[i].word
	}
	fmt.Println("wordlist",wordList)
	return resultList
}
func main() {
	strList := []string{"i", "love", "leetcode", "i", "love", "coding"}
	res := topKFrequent(strList, 2)
	fmt.Println("res", res)
}
