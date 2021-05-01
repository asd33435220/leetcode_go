package main

import "fmt"

func MaxInt(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func leastBricks(wall [][]int) int {
	res := 0

	myMap := make(map[int]int)
	for j := 0; j < len(wall); j++ {
		pos := 0
		for i := 0; i < len(wall[j])-1; i++ {
			pos += wall[j][i]
			myMap[pos]++
			res = MaxInt(res, myMap[pos])
		}
	}

	return len(wall) - res
}
func main() {
	//wall := [][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}}
	wall := [][]int{{7, 1, 2}, {3, 5, 1, 1}, {10}}
	res := leastBricks(wall)
	fmt.Println("res", res)
}
