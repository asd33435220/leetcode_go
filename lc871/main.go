package main

import (
	"container/heap"
	"fmt"
)

type hp struct {
	slice []int
}

func (h hp) Less(i, j int) bool {
	return h.slice[i] > h.slice[j]
}

func (h *hp) Push(v interface{}) {
	h.slice = append(h.slice, v.(int))
}
func (h *hp) Pop() interface{} {
	n := len(h.slice)
	newSlice, res := h.slice[:n-1], h.slice[n-1]
	h.slice = newSlice
	return res
}
func (h *hp) Swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
}
func (h hp) Len() int {
	return len(h.slice)
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	h := &hp{}
	heap.Init(h)
	ans := 0
	cur := 0
	for cur < len(stations) {
		if target <= startFuel {
			break
		}
		for h.Len() > 0 && startFuel < stations[cur][0] {
			startFuel += heap.Pop(h).(int)
			ans++
		}
		for i := cur; i < len(stations) && startFuel >= stations[i][0]; i++ {
			heap.Push(h, stations[i][1])

			cur = i + 1
		}
		if h.Len() == 0 {
			return -1
		}
	}
	for h.Len() > 0 && startFuel < target {
		startFuel += heap.Pop(h).(int)
		ans++
	}
	if startFuel < target {
		return -1
	}
	return ans
}
func main() {
	stations := [][]int{{25, 25}, {50, 50}}
	res := minRefuelStops(100, 50, stations)
	fmt.Println("res", res)
}
