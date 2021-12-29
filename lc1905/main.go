package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 3, 5}
	res := countQuadruplets(nums)
	fmt.Println("Res", res)

}

func countQuadruplets(nums []int) (res int) {
	// a+b+c+d => a+b = d-c
	myMap := make(map[int]int)
	for b := len(nums) - 3; b >= 1; b-- {
		for d := b + 2; d < len(nums); d++ {
			myMap[nums[d]-nums[b+1]]++
		}
		for a := 0; a < b; a++ {
			res += myMap[nums[a]+nums[b]]
		}

	}
	return
}
