package main

import (
	"fmt"
)

type appleHeap [][]int

func main() {
	rectangles := [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}
	res := countGoodRectangles(rectangles)
	fmt.Println("Res", res)

}

func countGoodRectangles(rectangles [][]int) int {
	max := 0
	maxList := []int{}
	for i := 0; i < len(rectangles); i++ {
		curr := getMax(rectangles[i])
		maxList = append(maxList, curr)
		if curr > max {
			max = curr
		}
	}
	count := 0
	for i := 0; i < len(maxList); i++ {
		if maxList[i] == max {
			count++
		}
	}
	return count
}

func getMax(rectangle []int) int {
	if rectangle[0] > rectangle[1] {
		return rectangle[1]
	} else {
		return rectangle[0]
	}
}
