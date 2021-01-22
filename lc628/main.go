package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 56, 74, 12, 4}
	res := maximumProduct(nums)
	fmt.Println(res)
}

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	fmt.Println("nums=",nums)
	if nums[0]*nums[1]*nums[len(nums)-1] > nums[len(nums)-3]*nums[len(nums)-2]*nums[len(nums)-1] {
		return nums[0]*nums[1]*nums[len(nums)-1]
	}else{
		return nums[len(nums)-3]*nums[len(nums)-2]*nums[len(nums)-1]
	}
}
