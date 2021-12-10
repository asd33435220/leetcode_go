package main

import (
	"fmt"
	"math"
)

func shortestCompletingWord(licensePlate string, words []string) (ans string) {
	resultList := make([]string, 0)
	myMap := make(map[string]int)
	for _, value := range licensePlate {
		if value < 64 {
			continue
		}
		if value > 64 && value < 91 {
			value += 32
		}
		myMap[string(value)]++
	}

	for _, word := range words {
		testMap := CopyMap(myMap)
		for _, value := range word {
			str := string(value)
			if _, ok := testMap[str]; ok {
				testMap[str]--
			}
		}
		flag := true
		for _, left := range testMap {
			if left > 0 {
				flag = false
			}
		}
		if flag {
			resultList = append(resultList, word)
		}
	}
	index := 0
	minLen := math.MaxInt64
	for key, value := range resultList {
		if len(value) < minLen {
			index = key
			minLen = len(value)
		}
	}

	return resultList[index]
}

func CopyMap(sourceMap map[string]int) map[string]int {
	newMap := make(map[string]int)
	for key, value := range sourceMap {
		newMap[key] = value
	}
	return newMap
}
func main() {
	licensePlate := "1s3 456"
	words := []string{"looks", "pest", "stew", "show"}
	res := shortestCompletingWord(licensePlate, words)
	fmt.Println("res", res)
}
