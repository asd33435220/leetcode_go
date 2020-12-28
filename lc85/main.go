package main

import "fmt"

func main() {
	matrix := [][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'}}
	res := maximalRectangle(matrix)
	fmt.Println("res=", res)
}
func maximalRectangle(matrix [][]byte) int {
	if len(matrix)==0 {
		return 0
	}
	arr := getRectangleList(matrix)
	res := 0
	for _, levelArr := range arr {
		levelRes := largestRectangleArea(levelArr)
		res = Max(res,levelRes)
	}
	return res
}
func getRectangleList(matrix [][]byte) [][]int {
	res := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		levelRes := make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				levelRes[j] = 0
			} else {
				height := 1
				for n := 1; i-n >= 0; n++ {
					if matrix[i-n][j] == '1' {
						height++
						continue
					} else {
						break
					}
				}
				levelRes[j] = height
			}
		}
		res[i] = levelRes
	}
	return res
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
