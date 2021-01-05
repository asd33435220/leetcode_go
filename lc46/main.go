package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)
}

func permute(nums []int) [][]int {
	var result [][]int
	var walk func([]int, int, int)
	walk = func(array []int, first int, length int) {
		if length == first {
			newArr := make([]int, len(array))
			copy(newArr, array)
			result = append(result, newArr)
		}
		for i := first; i < length; i++ {
			array = swap(array, i, first)
			walk(array, first+1, length)
			array = swap(array, i, first)
		}
	}
	walk(nums, 0, len(nums))
	return result
}
func swap(array []int, a int, b int) []int {
	tem := array[a]
	array[a] = array[b]
	array[b] = tem
	return array
}
