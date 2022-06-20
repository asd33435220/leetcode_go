package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	var walk func(node *TreeNode) string
	walk = func(node *TreeNode) string {
		if node == nil {
			return ""
		} else if node.Left == nil && node.Right == nil {
			return strconv.Itoa(node.Val)
		} else if node.Left == nil && node.Right != nil {
			return fmt.Sprintf("%d()(%s)", node.Val, walk(node.Right))
		} else if node.Left != nil && node.Right == nil {
			return fmt.Sprintf("%d(%s)", node.Val, walk(node.Left))
		} else {
			return fmt.Sprintf("%d(%s)(%s)", node.Val, walk(node.Left), walk(node.Right))
		}
	}
	result := walk(root)
	return result
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	res := tree2str(root)
	fmt.Println("res", res)
}
