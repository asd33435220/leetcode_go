package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type TreeNodeWithParentList struct {
	Val        int
	Left       *TreeNode
	Right      *TreeNode
	ParentList []*TreeNode
}

func indexOf(slice []*TreeNode, target *TreeNode) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	res := make([]int, 0)
	myMap := make(map[*TreeNode][]*TreeNode)
	var makeMap func(*TreeNode)
	makeMap = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			myMap[node.Left] = make([]*TreeNode, 0)
			myMap[node.Left] = append(myMap[node], node)
		}
		if node.Right != nil {
			myMap[node.Right] = make([]*TreeNode, 0)
			myMap[node.Right] = append(myMap[node], node)
		}

		if node == target {
			return
		}
		makeMap(node.Left)
		makeMap(node.Right)
	}

	makeMap(root)
	var getValue func(*TreeNode, int)
	getValue = func(node *TreeNode, count int) {
		if node == nil || node == target {
			return
		}
		if count == 0 && !indexOf(myMap[target], node) {
			res = append(res, node.Val)
			return
		}
		if !indexOf(myMap[target], node.Left) {
			getValue(node.Left, count-1)
		}
		if !indexOf(myMap[target], node.Right) {
			getValue(node.Right, count-1)
		}
	}

	getValue(target.Left, k-1)
	getValue(target.Right, k-1)

	for i := 0; i < k && len(myMap[target])-1-i >= 0; i++ {
		getValue(myMap[target][len(myMap[target])-1-i], k-1-i)
	}
	if k > 0 {
		if len(myMap[target]) >= k {
			res = append(res, myMap[target][len(myMap[target])-k].Val)
		}
	} else {
		res = append(res, target.Val)
	}

	return res
}
func main() {
	node := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:  7,
					Left: nil,
					Right: &TreeNode{
						Val:   8,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:  9,
						Left: nil,
						Right: &TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
		},
	}
	res := distanceK(node, node.Right.Right.Left, 5)
	fmt.Println("Res", res)
}
