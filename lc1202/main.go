package main

func main() {

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
func smallestStringWithSwaps(s string, pairs [][]int) string {
	//var Union []int
	//var res string
	//var getUnion func()
	//visited := make(map[int]bool)
	//groupNumber := 0
	//getUnion = func() {
	//	for key, value := range s {
	//
	//	}
	//}
	//getUnion()
	//
	//return res
}
