package main

import (
	"container/heap"
	"fmt"
)

type appleHeap [][]int

func (h appleHeap) Len() int           { return len(h) }
func (h appleHeap) Top() []int         { return h[0] }
func (h appleHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h appleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *appleHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *appleHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func main() {
	apples := []int{5, 2, 3}
	days := []int{6, 9, 10}
	res := eatenApples(apples, days)
	fmt.Println("Res", res)

}

func eatenApples(apples []int, days []int) int {
	myHeap := &appleHeap{}
	heap.Init(myHeap)
	ans := 0
	i := 0
	for i < len(apples) {
		for myHeap.Len() > 0 && myHeap.Top()[0] <= i {
			heap.Pop(myHeap)
		}
		if apples[i] != 0 {
			heap.Push(myHeap, []int{days[i] + i, apples[i]})
		}

		if myHeap.Len() > 0 {
			top := myHeap.Top()
			ans++
			top[1]--
			if top[1] == 0 {
				heap.Pop(myHeap)
			}
		} else if days[i] != 0 {
			return ans
		}
		i++
	}
	for myHeap.Len() > 0 {
		for myHeap.Len() > 0 && myHeap.Top()[0] <= i {
			heap.Pop(myHeap)
		}
		if myHeap.Len() > 0 {
			top := myHeap.Top()
			ans += Min(top[0]-i, top[1])
			i += Min(top[0]-i, top[1])
			heap.Pop(myHeap)
		} else {
			return ans
		}

	}
	return ans
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
