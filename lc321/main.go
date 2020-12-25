package main

import "fmt"

func main() {
	nums1 := []int{6,7}
	nums2 := []int{6,0,4}
	k := 5
	//result := getMaxStack(nums1, k)
	result := maxNumber(nums1, nums2, k)
	fmt.Println(result)
	//result := maxNumber(nums1, nums2, k)
	//fmt.Println(result)
}

//
//func getMaxStack(array []int, k int) []int {
//	var stack []int
//	for index, value := range array {
//		for len(stack) > 0 && value > stack[len(stack)-1] && len(array)+len(stack) > k+index+1 {
//			stack = stack[:len(stack)-1]
//		}
//		if len(stack) < k {
//			stack = append(stack, value)
//		}
//	}
//	return stack
//}
//func lexicographicalLess(a, b []int) bool {
//	for i := 0; i < len(a) && i < len(b); i++ {
//		if a[i] != b[i] {
//			return a[i] < b[i]
//		}
//	}
//	return len(a) < len(b)
//}
//
//func merge(a, b []int) []int {
//	result := make([]int, len(a)+len(b))
//	for i := range result {
//		if lexicographicalLess(a, b) {
//			result[i], b = b[0], b[1:]
//		} else {
//			result[i], a = a[0], a[1:]
//		}
//	}
//	return result
//}
//func maxNumber(nums1, nums2 []int, k int) []int {
//	var res []int
//	start := 0
//	if k > len(nums2) {
//		start = k - len(nums2)
//	}
//	for i := start; i <= k && i <= len(nums1); i++ {
//		s1 := getMaxStack(nums1, i)
//		s2 := getMaxStack(nums2, k-i)
//		merged := merge(s1, s2)
//		if lexicographicalLess(res, merged) {
//			res = merged
//		}
//	}
//	return res
//}

func getMaxStack(arr []int, k int) (stack []int) {
	for index, value := range arr {
		for len(stack) > 0 && stack[len(stack)-1] < value && len(stack)+len(arr) >= k+1+index {
			stack = stack[:len(stack)-1]
		}
		if len(stack) < k {
			stack = append(stack, value)
		}
	}
	return
}
func isBigger(a []int, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
func merge(a []int, b []int) (result []int) {
	result = make([]int, len(a)+len(b))
	for i, _ := range result {
		if !isBigger(a, b) {
			result[i], a = a[0], a[1:]
		} else {
			result[i], b = b[0], b[1:]
		}
	}
	return
}
func maxNumber(nums1 []int, nums2 []int, k int) (result []int) {
	start := 0
	if k > len(nums2) {
		start = k - len(nums2)
	}
	for i := start; i <= k && i <= len(nums1); i++ {
		stack1 := getMaxStack(nums1, i)
		stack2 := getMaxStack(nums2, k-i)
		merged := merge(stack1,stack2)
		if isBigger(result,merged){
			result = merged
		}
	}
	return
}
