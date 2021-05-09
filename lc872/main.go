package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var arr1, arr2 []int
	var walk func(node *TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			arr1 = append(arr1, node.Val)
		}
		walk(node.Left)
		walk(node.Right)
	}
	walk(root1)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			arr2 = append(arr2, node.Val)
		}
		walk(node.Left)
		walk(node.Right)
	}
	walk(root2)
	if len(arr1) != len(arr2) {
		return false
	} else {
		for i := 0; i < len(arr1); i++ {
			if arr1[i] != arr2[i] {
				return false
			}
		}
		return true
	}

}
func main() {

}
