package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	var walk func(node *TreeNode) bool
	walk = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		cutLeft := walk(node.Left)
		cutRight := walk(node.Right)

		if cutLeft {
			node.Left = nil
		}
		if cutRight {
			node.Right = nil
		}
		return cutLeft && cutRight && node.Val == 0
	}
	res := walk(root)
	if res {
		return nil
	} else {
		return root
	}
}
func main() {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	res := pruneTree(root)
	fmt.Println(res)
}
