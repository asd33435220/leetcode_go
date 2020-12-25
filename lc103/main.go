package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
	res := zigzagLevelOrder(node)
	fmt.Println(res)
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		levelRes := []int{}
		levelQue := queue
		queue = nil
		for _, node := range levelQue {
			levelRes = append(levelRes, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		if level%2 == 1 {
			for i, n := 0, len(levelRes); i < n/2; i++ {
				levelRes[i], levelRes[n-1-i] = levelRes[n-1-i], levelRes[i]
			}
		}
		res = append(res, levelRes)

	}
	return res
}
