package main

import (
	"fmt"
)

func main() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{Val: 5},
					{Val: 6},
				},
			},
			{
				Val:      2,
				Children: nil,
			},
			{
				Val:      4,
				Children: nil,
			},
		},
	}
	res := levelOrder(root)
	fmt.Println(res)
}

type Node struct {
	Val      int
	Children []*Node
}

type Node2 struct {
	Val      int
	Level    int
	Children []*Node
}

func levelOrder(root *Node) [][]int {

	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*Node2, 0)
	queue = append(queue, &Node2{root.Val, 1, root.Children})
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.Level > len(res) {
			res = append(res, []int{curr.Val})
		} else {
			level := curr.Level - 1
			res[level] = append(res[level], curr.Val)
		}
		for _, child := range curr.Children {
			queue = append(queue, &Node2{child.Val, curr.Level + 1, child.Children})
		}
	}
	return res
}
