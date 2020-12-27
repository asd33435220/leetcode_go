package main

import (
	"fmt"
)

func main() {
	reg := []int{2, 1, 2}
	res := largestRectangleArea(reg)
	fmt.Println(res)
}

func largestRectangleArea(heights []int) int {
	res := 0
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	n := len(heights)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			right := i
			res = Max(res, (right-left-1)*heights[cur])
		}
		stack = append(stack, i)
	}
	return res
}
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
