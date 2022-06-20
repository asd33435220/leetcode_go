package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res := levelOrder(root)
	fmt.Println("res", res)
}

type TreeNodeWithLevel struct {
	Node  *TreeNode
	Level int
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	queue := make([]*TreeNodeWithLevel, 1)
	queue[0] = &TreeNodeWithLevel{
		Node:  root,
		Level: 0,
	}
	for len(queue) > 0 {
		thisNode := queue[0]
		queue = queue[1:]
		if len(result) <= thisNode.Level {
			arr := make([]int, 1)
			arr[0] = thisNode.Node.Val
			result = append(result, arr)
		} else {
			result[thisNode.Level] = append(result[thisNode.Level], thisNode.Node.Val)
		}
		if thisNode.Node.Left != nil {
			nextNode := &TreeNodeWithLevel{
				Node:  thisNode.Node.Left,
				Level: thisNode.Level + 1,
			}
			queue = append(queue, nextNode)
		}
		if thisNode.Node.Right != nil {
			nextNode := &TreeNodeWithLevel{
				Node:  thisNode.Node.Right,
				Level: thisNode.Level + 1,
			}
			queue = append(queue, nextNode)
		}

	}
	return result
}
