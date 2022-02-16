package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	res := permute(nums)
	fmt.Println(res)
}

func permute(nums []int) [][]int {
	var result [][]int
	var walk func([]int, int)
	n := len(nums)
	walk = func(array []int, first int) {
		if n == first {
			newArr := make([]int, len(array))
			copy(newArr, array)
			result = append(result, newArr)
		}
		for i := first; i < n; i++ {
			array[i], array[first] = array[first], array[i]
			walk(array, first+1)
			array[i], array[first] = array[first], array[i]
		}
	}
	walk(nums, 0)
	return result
}
