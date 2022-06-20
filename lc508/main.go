package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: -3},
	}
	res := findFrequentTreeSum(root)
	fmt.Println("res", res)
}
func findFrequentTreeSum(root *TreeNode) (res []int) {
	max := 0
	hash := make(map[int]int)
	var walk func(node *TreeNode) int
	walk = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val + walk(node.Left) + walk(node.Right)
		hash[sum]++
		if hash[sum] > max {
			max = hash[sum]
		}
		return sum
	}
	walk(root)
	for key, val := range hash {
		if val == max {
			res = append(res, key)
		}
	}
	return
}
