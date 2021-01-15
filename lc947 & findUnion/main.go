package main

import "fmt"

func main() {
	stones := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}
	res := removeStones(stones)
	fmt.Println(res)
}

func removeStones(stones [][]int) int {
	union := &Union{
		f:     nil,
		count: 0,
	}
	union.init()
	for i := 0; i < len(stones); i++ {
		union.unionSet(stones[i][0], stones[i][1]+10000)
	}
	return len(stones) - union.count
}

type Union struct {
	f     map[int]int
	count int
}

func (union *Union) init() {
	union.f = make(map[int]int)
	union.count = 0
}
func (union *Union) add(x int) {
	_, ok := union.f[x]
	if !ok {
		union.f[x] = x
		union.count++
	}
}
func (union *Union) find(x int) int {
	root := x
	value, _ := union.f[x]
	for value != root {
		root, _ = union.f[root]
		value, _ = union.f[root]
	}
	for x != root {
		originFather, _ := union.f[x]
		union.f[x] = root
		x = originFather
	}
	return root

}

func (union *Union) unionSet(x int, y int) {
	union.add(x)
	union.add(y)
	rootX := union.find(x)
	rootY := union.find(y)
	if rootX == rootY {
		return
	}
	union.f[rootX] = y
	union.count--

}
