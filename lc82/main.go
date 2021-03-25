package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	res := deleteDuplicates(head)
	for res!= nil {
		fmt.Println(res.Val)
		head = res.Next
	}
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	thisNode := &ListNode{
		Val:  -10086,
		Next: head,
	}
	result := thisNode
	var lastNode *ListNode
	var nextNode *ListNode
	for thisNode != nil {
		nextNode = thisNode.Next
		if nextNode!=nil && thisNode.Val == nextNode.Val {
			for nextNode!=nil && thisNode.Val == nextNode.Val {
				nextNode = nextNode.Next
			}
			lastNode.Next = nextNode
		}else {
			lastNode = thisNode
		}
		thisNode = nextNode
	}
	return result.Next
}
