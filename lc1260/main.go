package main

import (
	"fmt"
)

func main() {
	indices := [][]int{{3, 8, 1, 9},
		{19, 7, 2, 5},
		{4, 6, 11, 10},
		{12, 0, 21, 13}}
	res := shiftGrid(indices, 4)
	fmt.Println(res)
}
func shiftGrid(grid [][]int, k int) [][]int {
	row := len(grid)
	col := len(grid[0])
	ans := make([][]int, row)
	for i := range ans {
		ans[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			index1 := (i*col + j + k) % (row * col)
			ans[index1/col][index1%col] = grid[i][j]
		}
	}
	return ans
}
