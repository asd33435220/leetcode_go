package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	res := numIslands(grid)
	fmt.Println(res)
}
func numIslands(grid [][]byte) (res int) {
	opts := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(grid), len(grid[0])
	var walk func(x, y int)
	walk = func(x, y int) {
		if grid[x][y] != '1' {
			return
		}
		grid[x][y] = '0'
		for _, opt := range opts {
			nextX, nextY := x+opt[0], y+opt[1]
			if nextX >= 0 && nextX < m && nextY >= 0 && nextY < n {
				walk(nextX, nextY)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				walk(i, j)
			}
		}
	}
	return
}
