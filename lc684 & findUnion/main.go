package main

import "fmt"

func main() {
	edges := [][]int{[]int{1, 2}, []int{1, 3}, []int{2, 3}}
	res := findRedundantConnection(edges)
	fmt.Println(res)
}

type Union struct {
	father map[int]int
}

func (union *Union) init() {
	union.father = make(map[int]int)
}

func (union *Union) add(x int) {
	_, ok := union.father[x]
	if !ok {
		union.father[x] = -1
	}
}

func (union *Union) find(x int) int {
	value, _ := union.father[x]
	root := x
	for value != -1 {
		root, _ = union.father[root]
		value, _ = union.father[root]
	}
	for x != root {
		originFather,_ := union.father[x]
		union.father[x] = root
		x = originFather
	}
	return root
}
func (union *Union) isConnected(x, y int) bool {
	return union.find(x) == union.find(y)
}
func (union *Union) merge(x, y int) {
	fatherX := union.find(x)
	fatherY := union.find(y)
	if fatherX != fatherY {
		union.father[fatherX] = fatherY
	}
}

func findRedundantConnection(edges [][]int) []int {
	union := &Union{}
	union.init()
	for i := 1; i <= len(edges); i++ {
		union.add(i)
	}
	for i := 0; i < len(edges); i++ {
		if union.isConnected(edges[i][0], edges[i][1]) {
			return edges[i]
		} else {
			union.merge(edges[i][0], edges[i][1])
		}
	}
	return []int{}
}
