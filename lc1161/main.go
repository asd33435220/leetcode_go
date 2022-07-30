package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   -8,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
	}
	res := maxLevelSum(root)
	fmt.Println("res", res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	max, maxLevel := math.MinInt64, 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		size := len(queue)
		sum := 0
		level++
		for i := 0; i < size; i++ {
			currentNode := queue[i]
			sum += currentNode.Val
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
		if sum > max {
			maxLevel = level
			max = sum
		}
		queue = queue[size:]
	}
	return maxLevel
}
