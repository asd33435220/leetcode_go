package main

import "fmt"

func main() {
	slice := []int{1, 3, 7, -4, 5, 6}
	res := findLaskK(slice, 3)
	fmt.Println(res)
}

func findLaskK(arr []int, k int) int {
	var walk func(arr []int, left int, right int, k int) int
	walk = func(arr []int, left int, right int, k int) int {
		if left >= right || len(arr) <= 1 {
			return arr[left]
		}
		pivot := findPivot(arr, left, right)
		if pivot+1 == k {
			return arr[pivot]
		} else if pivot+1 < k {
			return walk(arr, pivot+1, right, k)
		} else {
			return walk(arr, left, pivot-1, k)
		}

	}
	return walk(arr, 0, len(arr)-1, k)

}
func findPivot(arr []int, left int, right int) int {
	pivot := arr[left]
	for left < right {
		for left < right && pivot <= arr[right] {
			right--
		}
		swap(arr, left, right)
		for left < right && pivot >= arr[left] {
			left++
		}
		swap(arr, left, right)
	}
	return left
}
func swap(arr []int, left int, right int) {
	temp := arr[left]
	arr[left] = arr[right]
	arr[right] = temp
}
