package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Solution struct {
	rects      [][]int
	thresholds []int
}

func Constructor(rects [][]int) Solution {
	areas := make([]int, len(rects))
	var areaSum int
	for key, rect := range rects {
		areaSum += (rect[2] - rect[0] + 1) * (rect[3] - rect[1] + 1)
		areas[key] = areaSum
	}

	return Solution{
		rects:      rects,
		thresholds: areas,
	}
}

func (this *Solution) Pick() []int {
	threshold := rand.Intn(this.thresholds[len(this.thresholds)-1])
	index := sort.SearchInts(this.thresholds, threshold+1)
	rect := this.rects[index]
	x := rand.Intn(rect[2]-rect[0]+1) + rect[0]
	y := rand.Intn(rect[3]-rect[1]+1) + rect[1]
	return []int{x, y}
}
func main() {
	rects := [][]int{{-2, -2, 1, 1}, {2, 2, 4, 6}}
	slu := Constructor(rects)
	res := slu.Pick()
	fmt.Println("res", res)
	res = slu.Pick()
	fmt.Println("res", res)
	res = slu.Pick()
	fmt.Println("res", res)
}
