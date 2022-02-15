package main

import (
	"fmt"
)

func main() {
	// [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
	matrix := [][]int{
		{1, 10, 4, 2},
		{9, 3, 8, 7},
		{15, 16, 17, 12}}
	res := luckyNumbers(matrix)
	fmt.Println("res", res)
}
func luckyNumbers(matrix [][]int) []int {
	rowMap := make(map[int]int)
	colMap := make(map[int]int)
	res := make([]int, 0)
	for rowIndex, row := range matrix {
		for colIndex, value := range row {
			if v, ok := rowMap[rowIndex]; !ok || v > value {
				rowMap[rowIndex] = value
			}
			if colMap[colIndex] < value {
				colMap[colIndex] = value
			}
		}
	}
	for rowIndex, row := range matrix {
		for colIndex, value := range row {
			if rowMap[rowIndex] == value && colMap[colIndex] == value {
				res = append(res, value)
			}
		}
	}
	return res
}
