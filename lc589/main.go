package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) (res []int) {

	var walk func(node *Node)
	walk = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for _, child := range node.Children {
			walk(child)
		}
	}
	walk(root)
	return
}
