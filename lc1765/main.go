package main

import (
	"fmt"
)

func main() {
	isWater := [][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}
	res := highestPeak(isWater)
	fmt.Println("res", res)
}

func highestPeak(isWater [][]int) (res [][]int) {
	var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m := len(isWater)
	n := len(isWater[0])
	waterList := make([][]int, 0)
	for i := 0; i < m; i++ {
		list := make([]int, n)
		res = append(res, list)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				waterList = append(waterList, []int{i, j})
				res[i][j] = 0
			} else {
				isWater[i][j] = -1
			}
		}
	}

	for len(waterList) > 0 {
		currPoint := waterList[0]
		waterList = waterList[1:]
		value := res[currPoint[0]][currPoint[1]]
		for _, dir := range dirs {
			x, y := currPoint[0]+dir[0], currPoint[1]+dir[1]
			if x >= 0 && y >= 0 && x < m && y < n && isWater[x][y] == -1 {
				res[x][y] = value + 1
				isWater[x][y] = 0
				waterList = append(waterList, []int{x, y})
			}
		}

	}
	return res
}
