package main

import (
	"fmt"
	"sort"
)

func main() {
	p1 := []int{0, 0}
	p2 := []int{1, 0}
	p3 := []int{0, 1}
	p4 := []int{1, 1}
	res := validSquare(p1, p2, p3, p4)
	fmt.Println("res", res)
}
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	lens := []int{getLen(p1, p2), getLen(p1, p3), getLen(p1, p4), getLen(p2, p3), getLen(p4, p2), getLen(p3, p4)}
	sort.Ints(lens)
	return lens[0] > 0 && lens[0] == lens[1] && lens[0] == lens[2] && lens[0] == lens[3] && lens[4] == lens[5]
}
func getLen(p1 []int, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}
