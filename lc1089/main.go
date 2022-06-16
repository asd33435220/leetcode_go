package main

import (
	"fmt"
)

func duplicateZeros(arr []int) {
	var slice []int
	for _, val := range arr {
		fmt.Println(val)
		if val == 0 {
			slice = append(slice, 0)
			slice = append(slice, 0)
		} else {
			slice = append(slice, val)
		}
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = slice[i]
	}
}
func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
}
