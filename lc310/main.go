package main

import (
	"fmt"
	"math"
)

func main() {
	edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}
	result := findMinHeightTrees(6, edges)
	fmt.Println("max Int")
	fmt.Println(result)
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	res := make([]int, 0)
	graph := make([][]int, n)
	in := make([]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
		in[x]++
		in[y]++
	}
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if in[i] == 1 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		length := len(queue)
		res = make([]int, 0)
		for i := 0; i < length; i++ {
			curr := queue[0]
			queue = queue[1:]
			res = append(res, curr)
			for _, v := range graph[curr] {
				in[v]--
				if in[v] == 1 {
					queue = append(queue, v)
				}
			}
		}
	}
	return res
}

// 此方法会超时
func findMinHeightTreesOld(n int, edges [][]int) []int {
	graph := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}
	getMaxDepth := func(root int) int {
		maxDepth := 0
		visited := make(map[int]bool)
		var dfs func(root, deepth int)
		dfs = func(root, deepth int) {
			if visited[root] {
				return
			}
			visited[root] = true
			if deepth > maxDepth {
				maxDepth = deepth
			}
			for _, path := range graph[root] {
				dfs(path, deepth+1)
			}
		}
		dfs(root, 0)
		return maxDepth
	}
	depthList := make([]int, 0)
	min := math.MaxInt64
	for i := 0; i < n; i++ {
		depth := getMaxDepth(i)
		if depth < min {
			depthList = []int{i}
			min = depth
		} else if depth == min {
			depthList = append(depthList, i)
		}
	}

	return depthList

}
