package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var walk func(nums []int) *TreeNode
	getMax := func(nums []int) (max, index int) {
		for k, v := range nums {
			if v > max {
				index = k
				max = v
			}
		}
		return
	}
	walk = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		} else {
			max, index := getMax(nums)
			return &TreeNode{
				Val:   max,
				Left:  walk(nums[:index]),
				Right: walk(nums[index+1:]),
			}
		}

	}
	return walk(nums)
}
func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	res := constructMaximumBinaryTree(nums)
	fmt.Println("Res", res)
}
