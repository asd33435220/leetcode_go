package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  15,
			Left: nil,
			Right: &TreeNode{
				Val:   18,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res := rangeSumBST(root, 7, 15)
	fmt.Println(res)
}
func rangeSumBST(root *TreeNode, low int, high int) int {
	res := 0
	var walk func(*TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val <= high && node.Val >= low {
			res += node.Val
		}
		walk(node.Left)
		walk(node.Right)
	}
	walk(root)
	return res
}
