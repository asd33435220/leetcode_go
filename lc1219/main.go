package main

import "fmt"

var opts = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func getMaximumGold(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	ans := 0
	var dfs func(x, y, gold int)
	dfs = func(x, y, gold int) {
		gold += grid[x][y]
		if gold > ans {
			ans = gold
		}
		rec := grid[x][y]
		grid[x][y] = 0
		for _, opt := range opts {
			newX, newY := x+opt[0], y+opt[1]
			if newX < m && newY < n && newX >= 0 && newY >= 0 && grid[newX][newY] != 0 {
				dfs(newX, newY, gold)
			}
		}
		grid[x][y] = rec

	}
	for x, row := range grid {
		for y, gold := range row {
			if gold == 0 {
				continue
			}
			dfs(x, y, 0)
		}
	}
	return ans
}
func main() {
	arr := [][]int{
		{0, 6, 0},
		{5, 8, 7},
		{0, 9, 0},
		// {1, 0, 7},
		// {2, 0, 6},
		// {3, 4, 5},
		// {0, 3, 0},
		// {9, 0, 20}
	}
	res := getMaximumGold(arr)
	fmt.Println("res", res)
}
