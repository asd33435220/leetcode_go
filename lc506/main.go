package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	words := []int{10, 3, 8, 9, 4}
	res := findRelativeRanks(words)
	fmt.Println("res", res)
}

func findRelativeRanks(score []int) []string {
	ans := make([]string, len(score))
	myMap := make(map[int]int)
	for i := 0; i < len(score); i++ {
		myMap[score[i]] = i
	}
	s := make([]int, len(score))
	copy(s, score)
	sort.Ints(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	for i := 0; i < len(s); i++ {
		index := myMap[s[i]]
		if i == 0 {
			ans[index] = "Gold Medal"
			continue
		} else if i == 1 {
			ans[index] = "Silver Medal"
			continue
		} else if i == 2 {
			ans[index] = "Bronze Medal"
			continue
		} else {
			ans[index] = strconv.Itoa(i + 1)
		}
	}
	return ans
}
