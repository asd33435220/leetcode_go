package main

import (
	"fmt"
	"sort"
)

func main() {
	age := []int{20, 30, 100, 110, 120}
	res := numFriendRequests(age)
	fmt.Println(res)
}

func numFriendRequests(ages []int) int {
	res := 0
	sort.Ints(ages)
	left, right := 0, 0
	for i := 0; i < len(ages); i++ {
		if ages[i] <= 14 {
			continue
		}
		for 2*ages[left] <= ages[i]+14 {
			left++
		}
		for right+1 < len(ages) && ages[i] >= ages[right+1] {
			right++
		}
		res += right - left
	}
	return res
}
