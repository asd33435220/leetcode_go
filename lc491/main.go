package main

import (
	"fmt"
)

func nextGreaterElement(nums1 []int, nums2 []int) (res []int) {
	stack := make([]int, 0)
	resultMap := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 &&
			stack[len(stack)-1] < nums2[i] {
			resultMap[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	for i := 0; i < len(stack); i++ {
		resultMap[stack[i]] = -1
	}

	for i := 0; i < len(nums1); i++ {
		res = append(res, resultMap[nums1[i]])
	}
	return

}
func main() {
	nums1 := []int{137, 59, 92, 122, 52, 131, 79, 236, 94}
	nums2 := []int{137, 59, 92, 122, 52, 131, 79, 236, 94}
	res := nextGreaterElement(nums1, nums2)
	fmt.Println("res", res)
}
