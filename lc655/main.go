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

func printTree(root *TreeNode) [][]string {
	maxLevel := 0
	var getMaxLevel func(root *TreeNode, level int)
	getMaxLevel = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level > maxLevel {
			maxLevel = level
		}
		getMaxLevel(root.Left, level+1)
		getMaxLevel(root.Right, level+1)
	}
	getMaxLevel(root, 1)
	res := make([][]string, maxLevel)
	for i := 0; i < len(res); i++ {
		res[i] = make([]string, 1<<maxLevel-1)
	}
	var dfs func(node *TreeNode, x, y int)
	dfs = func(node *TreeNode, x, y int) {
		res[x][y] = strconv.Itoa(node.Val)
		if node.Left != nil {
			dfs(node.Left, x+1, y-1<<(maxLevel-1-x-1))
		}
		if node.Right != nil {
			dfs(node.Right, x+1, y+1<<(maxLevel-1-x-1))
		}
	}

	dfs(root, 0, (1<<maxLevel-1)/2)
	return res
}
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
	res := printTree(node)
	fmt.Println(res)
}
