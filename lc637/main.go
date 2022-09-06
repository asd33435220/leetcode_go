package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) (res []float64) {
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		list := queue
		queue = []*TreeNode{}
		sum := 0.0
		n := len(list)
		for i := 0; i < n; i++ {
			curr := list[i]
			sum += float64(curr.Val)
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		res = append(res, sum/float64(n))
	}
	return
}
func main() {

}
