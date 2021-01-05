package main

import "fmt"

func main() {
	//head := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 2,
	//				Next: &ListNode{
	//					Val: 5,
	//					Next: &ListNode{
	//						Val:  2,
	//						Next: nil,
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	res := partition(head, 2)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHead := large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next

}
