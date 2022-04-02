package main

import (
	"fmt"
)

func main() {
	candidates := []int{4, 2, 0, 3, 2, 5}
	res := trap(candidates)
	fmt.Println("res", res)
}

type Pillar struct {
	index  int
	height int
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func countWater(left, bottom, right *Pillar) int {
	width := right.index - left.index - 1
	height := min(right.height, left.height) - bottom.height
	return width * height
}
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	stack := make([]*Pillar, 0)
	ans := 0
	for key, value := range height {
		right := &Pillar{index: key, height: value}
		for len(stack) > 0 && stack[len(stack)-1].height <= value {
			bottom := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := stack[len(stack)-1]
				ans += countWater(left, bottom, right)
			}
		}
		stack = append(stack, right)
	}
	return ans
}
