package main

import "fmt"

func main() {
	arr := []int{1, 1, 2}
	res := permuteUnique(arr)
	fmt.Println(res)
}

func permuteUnique(nums []int) [][]int {
	var result [][]int
	var walk func(arr []int, first int, length int)
	walk = func(arr []int, first int, length int) {
		if first == length {
			newArr := make([]int,length)
			copy(newArr,arr)
			result = append(result, newArr)
			return
		}
		myMap := make(map[int]int)
		for i := first; i < length; i++ {
			value, ok := myMap[arr[i]]
			if ok || value > 0 {
				continue
			}
			myMap[arr[i]] = 1
			arr[first], arr[i] = arr[i], arr[first]
			walk(arr, first+1, length)
			arr[first], arr[i] = arr[i], arr[first]
		}

	}
	walk(nums, 0, len(nums))
	return result
}
