package main

import "fmt"

func lastK(arr []int, k int) {
	laskKEntry(arr, 0, len(arr)-1, k)
}
func laskKEntry(arr []int, left, right, k int) {
	if len(arr) <= 1 || left >= right {
		return
	}
	pivot := findPiovt(arr, left, right)
	if pivot == k {
		return
	} else if pivot > k {
		laskKEntry(arr, left, pivot-1, k)
	} else {
		laskKEntry(arr, pivot+1, right, k)
	}
}
func findPiovt(arr []int, left, right int) int {
	pivot := arr[left]
	for left < right {
		for left < right && pivot <= arr[right] {
			right--
		}
		arr[left],arr[right] = arr[right],arr[left]
		for left < right && pivot >= arr[left] {
			left++
		}
		arr[left],arr[right] = arr[right],arr[left]
	}
	return left
}

func smallestK(arr []int, k int) []int {
	lastK(arr, k)
	return arr[:k]
}
func main() {
	arr := []int{1, 3, 5, 7, 2, 4, 6, 8}
	res := smallestK(arr, 4)
	fmt.Println("res", res)
}
