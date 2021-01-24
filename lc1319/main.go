package main

import "fmt"

func main() {
	n := 6
	connections := [][]int{{0, 1},
		{0, 2},
		{0, 3},
		{1, 3},
		{1, 2}}
	res := makeConnected(n, connections)
	fmt.Println(res)
}

type FindUnion struct {
	f          map[int]int
	count      int // 联通分量
	redundancy int //冗余量
}

func (fu *FindUnion) init(n int) {
	fu.count = n
	fu.f = make(map[int]int)
	for i := 0; i < n; i++ {
		fu.f[i] = i
	}
}
func (fu *FindUnion) find(n int) int {
	if n != fu.f[n] {
		fu.f[n] = fu.find(fu.f[n])
	}
	return fu.f[n]
}
func (fu *FindUnion) merge(x, y int) {
	rootX := fu.find(x)
	rootY := fu.find(y)
	if rootX != rootY {
		fu.f[rootY] = rootX
		fu.count--
	} else {
		fu.redundancy++
	}
}

func makeConnected(n int, connections [][]int) int {
	fn := &FindUnion{}
	fn.init(n)
	for i := 0; i < len(connections); i++ {
		fn.merge(connections[i][0], connections[i][1])
	}
	if fn.redundancy < fn.count-1 {
		return -1
	} else {
		return fn.count - 1
	}

	//redundancy := 0
	//connectNumber := 0
	//used := make([]bool, len(connections))
	//visited := make([]bool, n)
	//var dfs func(node int)
	//dfs = func(node int) {
	//	visited[node] = true
	//	connectNumber++
	//	for i := 0; i < len(connections); i++ {
	//		if !used[i] && visited[connections[i][0]] && visited[connections[i][1]] {
	//			redundancy++
	//			used[i] = true
	//			continue
	//		}
	//		if connections[i][0] == node && !visited[connections[i][1]] {
	//			used[i] = true
	//			dfs(connections[i][1])
	//		} else if connections[i][1] == node && !visited[connections[i][0]] {
	//			used[i] = true
	//			dfs(connections[i][0])
	//		}
	//	}
	//}
	//dfs(connections[0][0])
	//if n-connectNumber <= redundancy {
	//	return n - connectNumber
	//} else {
	//	return -1
	//}
}
