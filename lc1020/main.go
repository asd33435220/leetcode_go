package main

import (
	"fmt"
	"strconv"
	"strings"
)

func numEnclaves(grid [][]int) (ans int) {
	opts := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	m := len(grid)
	n := len(grid[0])
	// pointMap := make([][]Point, m)
	// for i := 0; i < m; i++ {
	// 	pointMap[i] = make([]Point, n)
	// }
	// 1 未访问的
	// 2 已经访问的
	points := make(map[string]int)
	for y, rows := range grid {
		for x, value := range rows {
			if value == 1 {
				points[makePoint(x, y)] = 1
			}
		}
	}
	for key, value := range points {
		if value == 2 {
			continue
		}
		canLand := false
		queue := make([]string, 0)
		queue = append(queue, key)
		count := 0
		for len(queue) > 0 {
			thisPoint := queue[0]
			queue = queue[1:]
			if points[thisPoint] != 1 {
				continue
			}
			count++
			points[thisPoint] = 2
			x, y := getPoint(thisPoint)
			for _, opt := range opts {
				if !canLand && (x+opt[0] >= n || x+opt[0] == -1 || y+opt[1] >= m || y+opt[1] == -1) {
					canLand = true
				}
				nextPoint := makePoint(x+opt[0], y+opt[1])
				if points[nextPoint] == 1 {
					queue = append(queue, nextPoint)
				}
			}
		}
		if !canLand {
			ans += count
		}
	}
	return
}
func getPoint(point string) (int, int) {
	list := strings.Split(point, "-")
	x, _ := strconv.Atoi(list[0])
	y, _ := strconv.Atoi(list[1])
	return x, y
}
func makePoint(x, y int) string {
	return strconv.Itoa(x) + "-" + strconv.Itoa(y)
}
func main() {
	grid := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res := numEnclaves(grid)
	fmt.Println(res)
}
