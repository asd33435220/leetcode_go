package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	result := buildTree(preorder, inorder)
	fmt.Println(result)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var walk func([]int, []int) *TreeNode
	walk = func(preorder []int, inorder []int) *TreeNode {
		rootVal := preorder[0]
		root := &TreeNode{
			Val:   rootVal,
			Left:  nil,
			Right: nil,
		}
		inorderIndex := getIndex(inorder, rootVal)
		preorderIndex := getIndex(preorder, inorder[inorderIndex-1])
		leftPreorder := preorder[1:preorderIndex]
		rightPreorder := preorder[preorderIndex:]
		leftInorder := inorder[:inorderIndex]
		rightInorder := inorder[inorderIndex+1:]
		root.Left = walk(leftPreorder, leftInorder)
		root.Right = walk(rightPreorder, rightInorder)
		return root
	}
	res := walk(preorder, inorder)
	return res
}
func getIndex(intArray []int, target int) (index int) {
	for key, value := range intArray {
		if value == target {
			return key
		}
	}
	return -1
}
