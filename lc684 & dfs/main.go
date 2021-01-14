package main

import "fmt"

func main() {
	edges := [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}, {1, 5}}
	res := findRedundantConnection(edges)
	fmt.Println(res)
}

func findRedundantConnection(edges [][]int) []int {
	for i := 1; i < len(edges); i++ {
		visited := make([]bool, len(edges)+1)
		res := dfs(edges, i, visited)
		if res {
			return edges[i]
		}
	}
	return []int{}
}

func dfs(edges [][]int, number int, visited []bool) bool {
	entrance, target := edges[number][0], edges[number][1]
	var walk func(entrance int)
	flag := false
	walk = func(node int) {
		visited[node] = true
		for i := 0; i < number; i++ {
			if edges[i][0] == node && !visited[edges[i][1]] {
				if edges[i][1] == target {
					flag = true
					return
				}
				walk(edges[i][1])
			} else if edges[i][1] == node && !visited[edges[i][0]] {
				if edges[i][0] == target {
					flag = true
					return
				}
				walk(edges[i][0])
			}
		}

	}
	walk(entrance)
	return flag
}
