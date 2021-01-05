package main

import "fmt"

func main() {
	s := "aaa"
	res := largeGroupPositions(s)
	fmt.Println(res)
}

func largeGroupPositions(s string) [][]int {

	var array [][]int
	s += "0"
	lastChar := s[0]
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == lastChar {
			count++
		} else {
			if count >= 3 {
				ran := make([]int, 2)
				ran[0] = i - count
				ran[1] = i - 1
				array = append(array, ran)
			}
			count = 1
		}
		lastChar = s[i]
	}
	return array
}
