package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) (ans int) {
	var walk func(node *TreeNode) int
	walk = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sumLeft := walk(node.Left)
		sumRight := walk(node.Right)
		if sumRight > sumLeft {
			ans += sumRight - sumLeft
		} else {
			ans += sumLeft - sumRight
		}
		return sumLeft + sumRight + node.Val
	}
	walk(root)
	return

}
func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  8,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	res := findTilt(root2)
	fmt.Println("res", res)
	fmt.Println("res", root)

}
