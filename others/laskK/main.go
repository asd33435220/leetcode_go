package main

import "fmt"

func main() {
	slice := []int{4, 5, 1, 6, 2, 7, 3, 8}
	res := GetLeastNumbers_Solution(slice, 4)
	fmt.Println(res)
}
func GetLeastNumbers_Solution(input []int, k int) []int {
	// write code here
	return GetLastK(input, 0, len(input)-1, 4)
}
func GetLastK(input []int, left, right, k int) []int {

	pivot := findPivot(input, left, right)
	if pivot == k {
		QuickSort(input[:k], 0, k-1)
		return input[:k]
	} else if pivot > k {
		if pivot+1 == k {
			QuickSort(input[:k+1], 0, k)
			return input[:k+1]
		}
		GetLastK(input, left, pivot-1, k)
	} else {
		if pivot-1 == k {
			QuickSort(input[:k-1], 0, k-2)
			return input[:k-1]
		}
		GetLastK(input, pivot+1, right, k)
	}
	return []int{}
}
func QuickSort(arr []int, left, right int) {
	if len(arr) <= 1 || left >= right {
		return
	}
	pivot := findPivot(arr, left, right)
	QuickSort(arr, left, pivot-1)
	QuickSort(arr, pivot+1, right)
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
