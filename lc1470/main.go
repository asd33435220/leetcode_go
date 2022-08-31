package main

import "fmt"

func shuffle(nums []int, n int) (res []int) {
	left := nums[:n]
	right := nums[n:]
	for i := 0; i < 2*n; i++ {
		if i%2 == 0 {
			res = append(res, left[i/2])
		} else {
			res = append(res, right[i/2])
		}
	}
	return
}
func main() {
	target := []int{2, 5, 1, 3, 4, 7}
	fmt.Println("res", shuffle(target, 3))
}
