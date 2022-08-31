package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	var lastNode *TreeNode
	curr := root
	for curr != nil {
		if curr.Val < val {
			newRoot := &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			if lastNode == nil {
				newRoot.Left = curr
				return newRoot
			} else {
				lastNode.Right = newRoot
				newRoot.Left = curr
				return root
			}
		} else {
			lastNode = curr
			curr = curr.Right
		}
	}
	lastNode.Right = &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	return root
}
func main() {
	node := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}
	res := insertIntoMaxTree(node, 3)
	fmt.Println("Res", res)
}
