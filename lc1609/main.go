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

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res := isEvenOddTree(root)
	fmt.Println("res", res)
}

func isEvenOddTree(root *TreeNode) bool {
	// 递归遍历
	//treeArr := make([][]int, 0)
	//var walk func(node *TreeNode, level int)
	//walk = func(node *TreeNode, level int) {
	//	if node == nil {
	//		return
	//	}
	//	if len(treeArr) <= level {
	//		levelArr := make([]int, 0)
	//		levelArr = append(levelArr, node.Val)
	//		treeArr = append(treeArr, levelArr)
	//	} else {
	//		levelArr := treeArr[level]
	//		levelArr = append(levelArr, node.Val)
	//		treeArr[level] = levelArr
	//	}
	//	walk(node.Left, level+1)
	//	walk(node.Right, level+1)
	//}
	//walk(root, 0)
	//fmt.Println(treeArr)
	//for i := 0; i < len(treeArr); i++ {
	//	var last int
	//	if i % 2 == 0 {
	//		last = 0
	//		for j := 0; j < len(treeArr[i]); j++ {
	//			if treeArr[i][j] % 2 == 0 || treeArr[i][j] <= last {
	//				return false
	//			}
	//			last = treeArr[i][j]
	//		}
	//	}else{
	//		last = math.MaxInt64
	//		for j := 0; j < len(treeArr[i]); j++ {
	//			if treeArr[i][j] % 2 != 0 || treeArr[i][j] >= last {
	//				return false
	//			}
	//			last = treeArr[i][j]
	//		}
	//	}
	//
	//}
	//return true

	// 层序遍历
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		size := len(queue)
		var begin int
		if level%2 == 0 {
			begin = 0
		} else {
			begin = math.MaxInt64
		}
		for i := 0; i < size; i++ {
			if level%2 == 0 {
				currNode := queue[0]
				value := currNode.Val
				if value%2 == 0 || value <= begin {
					return false
				} else {
					begin = value
					queue = queue[1:]
					if currNode.Left != nil {
						queue = append(queue, currNode.Left)
					}
					if currNode.Right != nil {
						queue = append(queue, currNode.Right)
					}
				}
			} else {
				currNode := queue[0]
				value := currNode.Val
				if value%2 != 0 || value >= begin {
					return false
				} else {
					begin = value
					queue = queue[1:]
					if currNode.Left != nil {
						queue = append(queue, currNode.Left)
					}
					if currNode.Right != nil {
						queue = append(queue, currNode.Right)
					}
				}
			}
		}
		level++
	}
	return true
}
