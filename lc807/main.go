package main

import "fmt"

func main() {
	grid :=
		[][]int{
			{3, 0, 8, 4},
			{2, 4, 5, 7},
			{9, 2, 6, 3},
			{0, 3, 1, 0},
		}
	res := maxIncreaseKeepingSkyline(grid)
	fmt.Println(res)
}
func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	rowMax := make([]int, n)
	colMax := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rowMax[i] = Max(grid[i][j], rowMax[i])
			colMax[j] = Max(grid[i][j], colMax[j])
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			top := Min(colMax[j], rowMax[i])
			ans += top - grid[i][j]
		}
	}
	return ans
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func Min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
