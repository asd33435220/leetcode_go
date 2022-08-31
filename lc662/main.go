package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type MyNode struct {
	Node  *TreeNode
	Level int
}

func widthOfBinaryTree(root *TreeNode) int {
	queue := []*MyNode{
		{
			Node:  root,
			Level: 1,
		},
	}
	max := 0
	for len(queue) > 0 {
		list := queue
		n := len(list)
		queue = []*MyNode{}
		currWidth := list[n-1].Level - list[0].Level + 1
		if currWidth > max {
			max = currWidth
		}
		for i := 0; i < n; i++ {
			node := list[i].Node
			if node.Left != nil {
				queue = append(queue, &MyNode{Node: node.Left, Level: list[i].Level * 2})
			}
			if node.Right != nil {
				queue = append(queue, &MyNode{Node: node.Right, Level: list[i].Level*2 + 1})
			}
		}
	}
	return max
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
	res := widthOfBinaryTree(node)
	fmt.Println("Res", res)
}
