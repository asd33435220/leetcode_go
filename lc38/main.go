package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Same struct {
	str   string
	count int
}

func countSame(arr []string) []Same {
	arr = append(arr, "$")
	result := make([]Same, 0)
	lastStr := arr[0]
	count := 1
	for i := 1; i < len(arr); i++ {
		if lastStr == arr[i] {
			count++
		} else {
			result = append(result, Same{
				str:   lastStr,
				count: count,
			})
			count = 1
			lastStr = arr[i]
		}

	}
	return result
}

func countAndSay(n int) string {
	var walk func(n int) string
	walk = func(n int) string {
		if n == 1 {
			return "1"
		}
		result := ""
		arr := strings.Split(walk(n-1), "")
		sameArr := countSame(arr)
		for i := 0; i < len(sameArr); i++ {
			result += strconv.Itoa(sameArr[i].count) + sameArr[i].str
		}
		return result
	}
	return walk(n)
}
func main() {
	res := countAndSay(5)
	fmt.Println("res", res)
}
