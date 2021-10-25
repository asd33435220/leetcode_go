package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	arr := make([]int, 0)
	var walk func(node *TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		walk(node.Left)
		arr = append(arr, node.Val)
		walk(node.Right)
	}
	walk(root)
	return arr[len(arr)-k]
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}
	kthSmallest(root, 1)
}
