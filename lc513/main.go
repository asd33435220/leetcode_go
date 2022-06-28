package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) (res int) {
	curMax := -1
	var walk func(node *TreeNode, level int)
	walk = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level > curMax {
			curMax = level
			res = node.Val
		}
		walk(node.Left, level+1)
		walk(node.Right, level+1)
	}
	walk(root, 0)
	return
}

func main() {
	node := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	}
	res := findBottomLeftValue(node)
	fmt.Println("res", res)
}
