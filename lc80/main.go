package main

import "fmt"

func removeDuplicates(nums []int) int {
	lastNumber := nums[0]
	times := 1
	for i := 1; i < len(nums); i++ {
		if lastNumber == nums[i] {
			times++
			if times >= 3 {
				nums = append(nums[:i], nums[i+1:]...)
				i--
			}
		} else {
			lastNumber = nums[i]
			times = 1
		}
	}
	return len(nums)
}
func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	//nums2 := []int{0,1,2,3,4,5}
	//nums2 = append(nums2[:2],nums2[3:]...)
	//fmt.Println(nums2)
	res := removeDuplicates(nums)
	fmt.Println(nums)
	fmt.Println(res)
}
