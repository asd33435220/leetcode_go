package main

import (
	"fmt"
)

func main() {
	nums := "10"
	k := 1
	result := removeKdigits(nums, k)
	fmt.Println(result)
}

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	var stack []int32
	for _, char := range num {
		for k > 0 && len(stack) > 0 && char < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		if char != '0' || len(stack)!=0 {
			stack = append(stack,char)
		}
	}
	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	return string(stack)
}
