package main

import (
	"fmt"
)

func remove(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

type Tree struct {
	Val  int
	Left *Tree
}

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for _, v := range asteroids {
		alive := true
		for alive && v < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			if stack[len(stack)-1] == -v {
				stack = stack[:len(stack)-1]
				alive = false
			} else if stack[len(stack)-1] > -v {
				alive = false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		if alive {
			stack = append(stack, v)
		}
	}
	return stack
	//i := 0
	//for i < len(asteroids) && asteroids[i] < 0 {
	//	i++
	//}
	//for i < len(asteroids) {
	//	if asteroids[i] < 0 {
	//		for i > 0 && i < len(asteroids) && asteroids[i-1]*asteroids[i] < 0 {
	//			if asteroids[i-1] < -asteroids[i] {
	//				asteroids = remove(asteroids, i-1)
	//				i--
	//			} else if asteroids[i-1] > -asteroids[i] {
	//				asteroids = remove(asteroids, i)
	//				i--
	//				break
	//			} else {
	//				asteroids = remove(asteroids, i)
	//				asteroids = remove(asteroids, i-1)
	//				i -= 2
	//				break
	//			}
	//		}
	//		i++
	//	} else {
	//		i++
	//	}
	//}
	//return asteroids
}
func main() {
	nums := []int{5, 10, -10}
	res := asteroidCollision(nums)
	fmt.Println(res)
}
