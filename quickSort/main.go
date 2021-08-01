package main

import "fmt"

func main() {
	arr := []int{1, 6, 5, 9, -1}
	quickSort(arr,0, len(arr)-1)
	fmt.Println("arr", arr)
}

func quickSort(arr []int, left, right int) {
	if left >= right || len(arr) <= 1 {
		return
	}
	pivot := findPivot(arr, left, right)
	quickSort(arr, left, pivot-1)
	quickSort(arr, pivot+1, right)
}
func findPivot(arr []int, left, right int) int {
	pivot := arr[left]
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		arr[left], arr[right] = arr[right], arr[left]
		for left < right && arr[left] <= pivot {
			left++
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	return left
}
