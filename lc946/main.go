package main

import "fmt"

func main() {

	push := []int{2, 1, 0}
	poped := []int{1, 2, 0}
	fmt.Println(validateStackSequences(push, poped))
}

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	if n != len(popped) {
		return false
	}
	stack := make([]int, 0)
	popIndex := 0
	for i := 0; i < n; i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[popIndex] {
			stack = stack[:len(stack)-1]
			popIndex++
		}
	}
	return len(stack) == 0
}
