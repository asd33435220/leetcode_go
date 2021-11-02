package main

import (
	"fmt"
)

type union struct {
	index int
	tem   int
}

func dailyTemperatures(temperatures []int) (res []int) {
	stack := make([]union, 0)
	resultMap := make(map[union]int)
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && stack[len(stack)-1].tem < temperatures[i] {
			resultMap[stack[len(stack)-1]] = i - stack[len(stack)-1].index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, union{
			index: i,
			tem:   temperatures[i],
		})
	}

	for i := 0; i < len(stack); i++ {
		resultMap[stack[i]] = 0
	}
	for i := 0; i < len(temperatures); i++ {
		res = append(res, resultMap[union{
			index: i,
			tem:   temperatures[i],
		}])
	}
	return
}
func main() {
	nums := []int{73, 74, 75, 71, 69, 72, 76, 73}
	res := dailyTemperatures(nums)
	fmt.Println(res)
}
