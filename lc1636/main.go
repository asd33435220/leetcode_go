package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(frequencySort([]int{1, 1, 2, 2, 2, 3}))
}

func frequencySort(nums []int) []int {
	hash := make(map[int]int)
	for _, v := range nums {
		hash[v]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if hash[nums[i]] == hash[nums[j]] {
			return nums[i] > nums[j]
		} else {
			return hash[nums[i]] < hash[nums[j]]
		}
	})
	fmt.Println("nums", nums)
	return nums
}
