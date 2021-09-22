package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	count := 0
	countList := head
	res := make([]*ListNode, 0)
	for countList != nil {
		count++
		countList = countList.Next
	}
	number := count / k
	rest := count % k

	return res
}
func main() {
	fmt.Println("hello world")
	list := &ListNode{
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
	res := splitListToParts(list, 3)
	fmt.Println("res", res)
}
