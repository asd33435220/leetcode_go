package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type CBTInserter struct {
	root           *TreeNode
	nextParentNode *TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	return CBTInserter{
		root:           root,
		nextParentNode: walk(root),
	}
}

func walk(node *TreeNode) (res *TreeNode) {
	que := make([]*TreeNode, 0)
	que = append(que, node)
	var currNode *TreeNode
	for len(que) > 0 {
		currNode = que[0]
		que = que[1:]
		if currNode.Left != nil && currNode.Right == nil {
			return currNode
		}
		if currNode.Left == nil && currNode.Right == nil {
			return currNode
		}
		que = append(que, currNode.Left, currNode.Right)
	}
	return currNode
}

func (this *CBTInserter) Insert(val int) int {
	res := this.nextParentNode.Val
	if this.nextParentNode.Left == nil {
		this.nextParentNode.Left = &TreeNode{
			Val: val,
		}
	} else {
		this.nextParentNode.Right = &TreeNode{
			Val: val,
		}
		this.nextParentNode = walk(this.root)
	}
	return res
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
