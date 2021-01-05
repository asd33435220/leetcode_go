package main

import "fmt"

func main() {
	s := "abb"
	res := permuteUnique(s)
	fmt.Println(res)
}

func permuteUnique(s string) []string {
	var result []string
	var walk func(arr []int32, first int, length int)
	walk = func(arr []int32, first int, length int) {
		if first == length {
			newArr := make([]int32, length)
			copy(newArr, arr)
			result = append(result, string(newArr))
			return
		}
		myMap := make(map[int32]int32)
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
	var charArr []int32
	for _,value := range s {
		charArr = append(charArr,value)
	}
	walk(charArr, 0, len(charArr))
	return result
}
