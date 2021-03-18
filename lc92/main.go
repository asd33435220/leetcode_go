package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	//head := &ListNode{
	//	Val:  1,
	//	Next: nil,
	//}
	res := reverseBetween(head, 2, 4)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var prevNode *ListNode = nil
	var nextNode *ListNode = nil
	thisNode := head
	i := 1
	for thisNode != nil {
		if i == left {
			listNode1 := prevNode
			listNode2 := thisNode
			prevNode = nil
			for i <= right {
				nextNode = thisNode.Next
				thisNode.Next = prevNode
				prevNode = thisNode
				thisNode = nextNode
				i++
			}
			if listNode1 == nil {
				head = prevNode
			} else {
				listNode1.Next = prevNode
			}
			listNode2.Next = nextNode
			return head
		}
		prevNode = thisNode
		thisNode = thisNode.Next
		i++
	}
	return head
}
