package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "this apple is sweet"
	str2 := "this apple is sour"
	res := uncommonFromSentences(str1, str2)
	fmt.Println(res)
}
func uncommonFromSentences(s1 string, s2 string) (res []string) {
	strList1 := strings.Split(s1, " ")
	strList2 := strings.Split(s2, " ")
	resultMap := make(map[string]int)
	for _, value := range strList1 {
		resultMap[value]++
	}
	for _, value := range strList2 {
		resultMap[value]++
	}
	for key, value := range resultMap {
		if value == 1 {
			res = append(res, key)
		}
	}
	return
}
