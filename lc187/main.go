package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	strMap := make(map[string]int)
	byteSlice := []byte(s)
	for i := 0; i < len(byteSlice)-9; i++ {
		str := byteSlice[i : i+10]
		strMap[string(str)]++
	}
	result := make([]string, 0)
	for key, value := range strMap {
		if value > 1 {
			result = append(result, key)
		}
	}
	return result
}
func main() {
	s := "AAAAAAAAAAA"
	res := findRepeatedDnaSequences(s)
	fmt.Println("res", res)
}
