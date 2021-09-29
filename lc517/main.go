package main

import "fmt"

func findMinMoves(machines []int) int {
	total := 0
	for _, v := range machines {
		total += v
	}
	n := len(machines)
	if total%n != 0 {
		return -1
	}
	sum := 0
	res := 0
	avg := total / n
	for _, num := range machines {
		num -= avg
		sum += num
		res = max(res, max(abs(sum), num))
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func main() {
	nums := []int{1, 0, 5}
	res := findMinMoves(nums)
	fmt.Println("res", res)
}
