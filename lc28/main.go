package main

import "fmt"

func getNext(str string) []int {
	next := make([]int, len(str))
	next[0] = 0
	k := 0
	for i := 1; i < len(str); i++ {
		for k > 0 && str[i] != str[k] {
			k = next[k-1]
		}
		if str[k] == str[i] {
			k++
		}
		next[i] = k
	}
	return next
}
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	next := getNext(needle)
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			flag := true
			j := 1
			i++
			for ; j < len(needle) && i < len(haystack); {
				if needle[j] == haystack[i] {
					j++
					i++
				} else {
					flag = false
					break
				}
			}
			if flag && j == len(needle) {
				return i - j
			} else {
				i = i - j + next[j]
			}
		}
	}
	return -1
}
func main() {
	res := strStr("mississippi", "issip")
	fmt.Println(res)
}
