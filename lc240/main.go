package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	height := len(matrix)
	if height <= 0 {
		return false
	}
	width := len(matrix[0])
	if width <= 0 {
		return false
	}
	x := width - 1
	y := 0
	for x >= 0 && y < height {
		num := matrix[y][x]
		if num == target {
			return true
		} else if num < target {
			y++
		} else {
			x--
		}
	}
	return false

}

func main() {
	martix := [][]int{
		{-5}, {6},
	}
	//{1, 4, 7, 11, 15},
	//{2, 5, 8, 12, 19},
	//{3, 6, 9, 16, 22},
	//{10, 13, 14, 17, 24},
	//{18, 21, 23, 26, 30}}
	res := searchMatrix(martix, 6)
	fmt.Println("res", res)
}
