package main

import "fmt"

func arrayNesting(nums []int) (res int) {
	visited := make(map[int]bool)
	var walk func(index, count int) int
	walk = func(index, count int) int {
		if visited[index] {
			return count
		} else {
			visited[index] = true
			return walk(nums[index], count+1)
		}
	}

	for i := range nums {
		v := walk(i, 0)
		if v > res {
			res = v
		}
	}
	return
}
func main() {

	res := arrayNesting([]int{5, 4, 0, 3, 1, 6, 2})
	fmt.Println("res", res)
}
