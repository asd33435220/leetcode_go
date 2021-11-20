package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	provincesNumber := findLHS(nums)
	fmt.Println(provincesNumber)
}

func findLHS(nums []int) (ans int) {
	myMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		myMap[nums[i]]++
	}

	for key, value := range myMap {
		if v := myMap[key+1]; v != 0 && v+value > ans {
			ans = v + value
		}
	}
	return
}
