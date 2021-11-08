package main

import "fmt"

func longestSubsequence(arr []int, difference int) (ans int) {
	myMap := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		_, ok := myMap[arr[i]-difference]
		if ok {
			myMap[arr[i]] = myMap[arr[i]-difference] + 1
			ans = max(myMap[arr[i]], ans)
		} else {
			myMap[arr[i]] = 1
		}
	}
	return
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func main() {
	arr := []int{1, 5, 7, 8, 5, 3, 4, 2, 1, -2}
	res := longestSubsequence(arr, -2)
	fmt.Println("res", res)
}
