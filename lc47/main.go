package main

import "fmt"

func main() {
	arr := []int{1, 1, 2}
	res := permuteUnique(arr)
	fmt.Println(res)
}

func permuteUnique(nums []int) [][]int {
	var result [][]int
	n := len(nums)
	var walk func(arr []int, first int)
	walk = func(arr []int, first int) {
		if first == n {
			newArr := make([]int, n)
			copy(newArr, arr)
			result = append(result, newArr)
			return
		}
		myMap := make(map[int]bool)
		for i := first; i < n; i++ {
			if value := myMap[arr[i]]; value {
				continue
			}
			myMap[arr[i]] = true
			arr[first], arr[i] = arr[i], arr[first]
			walk(arr, first+1)
			arr[first], arr[i] = arr[i], arr[first]
		}

	}
	walk(nums, 0)
	return result
}
