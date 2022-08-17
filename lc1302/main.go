package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res := deepestLeavesSum(node)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	sum := 0
	for len(queue) > 0 {
		node := append([]*TreeNode{}, queue...)
		queue = []*TreeNode{}
		sum = 0
		for i := 0; i < len(node); i++ {
			curr := node[i]
			sum += curr.Val
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}
	return sum
}
