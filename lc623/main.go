package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	level := 0
	for len(q) > 0 {
		size := len(q)
		level++
		if level+1 > depth {
			break
		}
		for i := 0; i < size; i++ {
			curr := q[0]
			q = q[1:]
			if level == depth-1 {
				left := curr.Left
				right := curr.Right
				curr.Left = &TreeNode{
					Val:   val,
					Left:  left,
					Right: nil,
				}
				curr.Right = &TreeNode{
					Val:   val,
					Left:  nil,
					Right: right,
				}
			} else {
				if curr.Left != nil {
					q = append(q, curr.Left)
				}
				if curr.Right != nil {
					q = append(q, curr.Right)
				}
			}
		}
	}
	return root
}
func main() {
	node := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res := addOneRow(node, 1, 2)
	fmt.Println("res", res)
}
