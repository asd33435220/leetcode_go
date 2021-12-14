package main

import (
	Heap "../heap"
	"fmt"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	myHeap := &Heap.Heap{}
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	sum := 0
	for i := 0; i < len(courses); i++ {
		myHeap.Push(courses[i][0])
		sum += courses[i][0]
		if sum > courses[i][1] {
			duration := myHeap.Pop()
			sum -= duration
		}
	}
	return len(myHeap.IntSlice)
}
func main() {
	courses := [][]int{{100, 200},
		{200, 1300},
		{1000, 1250},
		{2000, 3200}}
	res := scheduleCourse(courses)
	fmt.Println(res)
}
