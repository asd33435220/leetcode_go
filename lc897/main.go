package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	var walk func(*TreeNode)
	result := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	result2 := result
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		walk(node.Left)
		child := &TreeNode{
			Val:   node.Val,
			Left:  nil,
			Right: nil,
		}
		result.Right = child
		result = result.Right
		walk(node.Right)
	}
	walk(root)
	return result2.Right
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  6,
			Left: nil,
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	res := increasingBST(root)
	fmt.Println(res)
}
