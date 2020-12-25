package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{3, &ListNode{4, &ListNode{5, nil}}}
	arr := reversePrint(list)
	fmt.Println(arr)
}
func reversePrint(head *ListNode) []int {
	head = reverseList(head)
	var arr []int
	thisNode := head
	for thisNode!=nil {
		arr = append(arr,thisNode.Val)
		thisNode = thisNode.Next
	}
	return arr
}
func reverseList(head *ListNode) *ListNode {
	thisNode := head
	var prev *ListNode = nil
	var nextNode *ListNode = nil
	for thisNode!=nil {
		nextNode = thisNode.Next
		thisNode.Next = prev
		prev = thisNode
		thisNode = nextNode
	}
	return prev
}
