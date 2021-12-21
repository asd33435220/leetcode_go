package main

import (
	"fmt"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	ans := 0
	for i, j := 0, 0; i < len(houses); i++ {
		cur := ABS(houses[i] - heaters[j])
		for j < len(heaters) && heaters[j] <= houses[i] {
			cur = houses[i] - heaters[j]
			j++
		}
		if j < len(heaters) {
			if v := heaters[j] - houses[i]; v < cur {
				cur = v
			}
		}
		if cur > ans {
			ans = cur
		}
		if j > 0 {
			j--
		}
	}
	return ans

}

func ABS(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func main() {
	res := findRadius([]int{1, 7}, []int{3, 5})
	fmt.Println("res", res)
}
