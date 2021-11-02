package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	root := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	deleteNode(root)
	fmt.Println("root", root)
}
