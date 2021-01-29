package main

import "fmt"

func main() {
	n := 4
	edges := [][]int{{3, 1, 2},
		{3, 2, 3},
		{1, 1, 3},
		{1, 2, 4},
		{1, 1, 2},
		{2, 3, 4}}
	res := maxNumEdgesToRemove(n, edges)
	fmt.Println(res)
}

type FindUnion struct {
	father map[int]int
	count  int
}

func (this *FindUnion) init(n int) {
	this.father = make(map[int]int)
	this.count = n
	for i := 0; i <= n; i++ {
		this.father[i] = i
	}
}
func (this *FindUnion) find(x int) int {
	value, _ := this.father[x]
	if value != x {
		this.father[x] = this.find(this.father[x])
	}
	return this.father[x]
}
func (this *FindUnion) merge(x, y int) bool {
	rootX := this.find(x)
	rootY := this.find(y)
	if rootX == rootY {
		return false
	} else {
		this.father[rootY] = rootX
		this.count--
		return true
	}
}
func maxNumEdgesToRemove(n int, edges [][]int) int {
	res := 0
	fua := &FindUnion{}
	fua.init(n)
	fub := &FindUnion{}
	fub.init(n)
	for i := 0; i < len(edges); i++ {
		if edges[i][0] == 3 {
			if !fua.merge(edges[i][1], edges[i][2]) {
				res++
			} else {
				fub.merge(edges[i][1], edges[i][2])
			}
		}
	}
	for i := 0; i < len(edges); i++ {
		if edges[i][0] == 1 {
			if !fua.merge(edges[i][1], edges[i][2]) {
				res++
			}
		}
	}
	for i := 0; i < len(edges); i++ {
		if edges[i][0] == 2 {
			if !fub.merge(edges[i][1], edges[i][2]) {
				res++
			}
		}
	}
	if fua.count != 1 || fub.count != 1 {
		return -1
	}
	return res
}
