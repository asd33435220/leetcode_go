package main

import "fmt"

func main() {
	isConnected := [][]int{[]int{1, 0, 0, 1}, []int{0, 1, 1, 0}, []int{0, 1, 1, 1}, []int{1, 0, 1, 1}}
	provincesNumber := findCircleNum(isConnected)
	fmt.Println(provincesNumber)
}

func findCircleNum(isConnected [][]int) int {
	var bfs func([][]int, []bool, int, int)
	bfs = func(isConnected [][]int, visited []bool, citiesNumber int, i int) {
		var queue []int
		queue = append(queue, i)
		for len(queue) > 0 {
			thisCity := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			if visited[thisCity] {
				continue
			}
			visited[thisCity] = true

			for j := 0; j < citiesNumber; j++ {
				if isConnected[j][thisCity] == 1 && !visited[j] {
					queue = append(queue, j)
				}
			}
		}
	}
	provincesNumber := 0
	citiesNumber := len(isConnected[0])
	visited := make([]bool, citiesNumber)
	for i := 0; i < citiesNumber; i++ {
		visited[i] = false
	}
	for i := 0; i < citiesNumber; i++ {
		if visited[i] {
			continue
		}
		bfs(isConnected, visited, citiesNumber, i)
		provincesNumber++

	}

	return provincesNumber
}
