package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) (ans *TreeNode) {
	var walk func(node *TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val == val {
			ans = node
			return
		} else if node.Val > val {
			walk(node.Left)
		} else {
			walk(node.Right)
		}
	}
	walk(root)
	return
}
func main() {
}
