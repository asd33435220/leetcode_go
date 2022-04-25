package main

import "fmt"

func main() {
	//input := [][]int{
	//	{1, 4, 7, 11, 15},
	//	{2, 5, 8, 12, 19},
	//	{3, 6, 9, 16, 22},
	//	{10, 13, 14, 17, 24},
	//	{18, 21, 23, 26, 30}}
	input := [][]int{{0, 0}}
	result := findNumberIn2DArray(input, 20)
	fmt.Print(result)
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	colNumber := len(matrix)
	if colNumber == 0 {
		return false
	}
	rowNumber := len(matrix[0])
	if rowNumber == 0 {
		return false
	}
	col := 0
	row := rowNumber - 1
	for col < colNumber && row >= 0 {
		fmt.Println(row, col)
		number := matrix[col][row]
		if number == target {
			return true
		} else if number < target {
			col++
		} else {
			row--
		}
	}
	return false
}
