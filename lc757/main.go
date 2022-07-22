package main

import (
	"fmt"
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] > intervals[j][0]
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})
	// min1 次小值 min2最小值
	min1, min2, ans := -1, -1, 0

	for i := 0; i < len(intervals); i++ {
		cur := intervals[i]
		// pre, pos
		// [cur[0], cur[1]]
		if cur[0] <= min2 {
			continue // 全包
		}
		// cur[0] > pre
		if cur[0] <= min1 { // 压中一个
			ans++
			min2 = min1
			min1 = cur[1]
		} else { // cur[0] > pos , 没有交集
			ans += 2
			min1 = cur[1]
			min2 = min1 - 1
		}
	}
	return ans
}
func main() {
	arr := [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}}
	res := intersectionSizeTwo(arr)
	fmt.Println("res", res)
}
