package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) (res []int) {
	hash := make(map[int]int)
	var walk func(node *TreeNode, level int)
	maxLevel := -1
	walk = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level >= maxLevel {
			maxLevel = level
		}
		v, ok := hash[level]
		val := node.Val
		if !ok || v < val {
			hash[level] = val
		}
		walk(node.Left, level+1)
		walk(node.Right, level+1)
	}
	walk(root, 0)
	if maxLevel == -1 {
		return []int{}
	}
	res = make([]int, maxLevel+1)
	for key, val := range hash {
		res[key] = val
	}
	return
}

func main() {
}
