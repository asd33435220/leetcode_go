package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func Permutaion(str string) (res []string) {
	k := len(str)
	var walk func(s []byte, n int)
	walk = func(s []byte, n int) {
		if n == k {
			res = append(res, string(s))
			return
		}
		hash := make(map[byte]bool)
		for i := n; i < k; i++ {
			if hash[s[i]] {
				continue
			}
			hash[s[i]] = true
			s[i], s[n] = s[n], s[i]
			walk(s, n+1)
			s[i], s[n] = s[n], s[i]
		}
	}
	walk([]byte(str), 0)
	return
}

func nextGreaterElement(n int) int {
	str := strconv.Itoa(n)
	arr := Permutaion(str)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			if i == len(arr)-1 {
				return -1
			} else {
				res, _ := strconv.Atoi(arr[i+1])
				if res > math.MaxInt32 {
					return -1
				}
				return res
			}
		}
	}
	return -1
}
func main() {
	res := nextGreaterElement(124)
	fmt.Println("res", res)
}
