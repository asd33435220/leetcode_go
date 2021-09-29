package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	preSum := make(map[int]int)
	preSum[0] = 1
	var walk func(*TreeNode, int)
	walk = func(node *TreeNode, currSum int) {
		if node == nil {
			return
		}
		currSum += node.Val
		res += preSum[currSum-targetSum]
		preSum[currSum]++
		walk(node.Left, currSum)
		walk(node.Right, currSum)
		preSum[currSum]--
	}
	walk(root, 0)
	return res
}
func main() {
	node := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   -2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:  2,
				Left: nil,
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:  -3,
			Left: nil,
			Right: &TreeNode{
				Val:   11,
				Left:  nil,
				Right: nil,
			},
		},
	}
	pathSum(node, 8)
}
