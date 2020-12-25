package main

import (
	"fmt"
)

func main() {
	rat := []int{1, 2, 2}
	res := candy(rat)
	fmt.Println(res)
}

//func candy(ratings []int) int {
//	res := 0
//	leftArr := make([]int, len(ratings))
//	for i, _ := range ratings {
//		if i > 0 && ratings[i] > ratings[i-1] {
//			leftArr[i] = leftArr[i-1] + 1
//		} else {
//			leftArr[i] = 1
//		}
//	}
//	rightNum := 0
//	for i := len(ratings) - 1; i >= 0; i-- {
//		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
//			rightNum++
//		} else {
//			rightNum = 1
//		}
//
//		res += max(leftArr[i], rightNum)
//	}
//
//	return res
//}
//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}

func candy(ratings []int) int {
	n := len(ratings)
	ans, inc, dec, pre := 1, 1, 0, 1
	for i:=1;i<n;i++ {
		if ratings[i]>=ratings[i-1]{
			dec = 0
			if ratings[i]==ratings[i-1]{
				pre = 1
			}else{
				pre++
			}
			ans += pre
			inc = pre
		}else{
			dec++
			if inc == dec {
				dec++
			}
			ans += dec
			pre = 1
		}
	}
	return ans
}
