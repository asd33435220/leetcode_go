package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) (list []*TreeNode) {
	var walk func(node *TreeNode) string
	paths := make(map[string]int)
	walk = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		res := strconv.Itoa(node.Val) + "l" + walk(node.Left) + "r" + walk(node.Right)
		paths[res]++
		if paths[res] == 2 {
			list = append(list, node)
		}
		return res
	}
	walk(root)
	return
}
func main() {

}
