package main

import "fmt"

func canBeEqual(target []int, arr []int) bool {
	m := len(target)
	if m != len(arr) {
		return false
	}
	hash := make(map[int]int)
	for i := 0; i < m; i++ {
		hash[target[i]]++
		hash[arr[i]]--
	}
	for _, v := range hash {
		if v != 0 {
			return false
		}
	}
	return true
}
func main() {
	target := []int{1, 2, 3, 4}
	arr := []int{2, 4, 1, 3}
	fmt.Println("res", canBeEqual(target, arr))
}
