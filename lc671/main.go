package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if root.Left == nil {
		return -1
	}
	if root.Left.Val == root.Right.Val {
		return -1
	}
	ans := math.MaxInt64
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val < root.Val {
			return
		} else {
			if node.Val < ans && node.Val > root.Val {
				ans = node.Val
			}
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root.Left)
	dfs(root.Right)
	if ans == math.MaxInt64 {
		return  -1
	}
	return ans
}
func main() {
	node := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   5,
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
	res := findSecondMinimumValue(node)
	fmt.Println("Res", res)
}
