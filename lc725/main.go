package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	res := make([]*ListNode, k)
	count := 0
	countList := head
	for countList != nil {
		count++
		countList = countList.Next
	}
	number := count / k
	rest := count % k
	curr := head
	for i := 0; i < k && curr != nil; i++ {
		res[i] = curr
		size := number
		if rest > 0 {
			size++
		}
		for j := 0; j < size-1; j++ {
			curr = curr.Next
		}
		next := curr.Next
		curr.Next = nil
		curr = next
		rest--
	}

	return res
}
func main() {
	fmt.Println("hello world")
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
				// Next: &ListNode{
				// 	Val: 4,
				// 	Next: &ListNode{
				// 		Val: 5,
				// 		Next: &ListNode{
				// 			Val: 6,
				// 			Next: &ListNode{
				// 				Val: 7,
				// 				Next: &ListNode{
				// 					Val: 8,
				// 					Next: &ListNode{
				// 						Val:  9,
				// 						Next: nil,
				// 					},
				// 				},
				// 			},
				// 		},
				// 	},
				// },
			},
		},
	}
	res := splitListToParts(list, 2)
	fmt.Println("res", res)
}
