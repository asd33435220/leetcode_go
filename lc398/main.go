package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{1, 2, 3, 3, 3}
	slu := Constructor(nums)
	fmt.Println(slu.Pick(3))
	fmt.Println(slu.Pick(3))
	fmt.Println(slu.Pick(3))

}

type Solution struct {
	hash map[int][]int
}

func Constructor(nums []int) Solution {
	hash := make(map[int][]int)
	for key, value := range nums {
		hash[value] = append(hash[value], key)
	}
	return Solution{
		hash: hash,
	}
}

func (this *Solution) Pick(target int) int {
	n := len(this.hash[target])
	index := rand.Intn(n)
	return this.hash[target][index]
}
