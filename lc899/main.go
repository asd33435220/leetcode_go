package main

import (
	"fmt"
	"sort"
)

func orderlyQueue(s string, k int) string {
	if k > 1 {
		byteArr := []byte(s)
		sort.Slice(byteArr, func(i, j int) bool {
			return byteArr[i] < byteArr[j]
		})
		return string(byteArr)
	} else {
		n := len(s)
		strArr := make([]string, n)
		byteArr := []byte(s)
		for i := 0; i < n; i++ {
			front := byteArr[0]
			byteArr = byteArr[1:]
			byteArr = append(byteArr, front)
			strArr[i] = string(byteArr)
		}
		sort.Strings(strArr)
		return strArr[0]
	}
}
func main() {
	res := orderlyQueue("baaca", 3)
	fmt.Println(res)
}
