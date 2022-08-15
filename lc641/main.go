package main

import (
	"fmt"
)

type MyCircularDeque struct {
	queue []int
	size  int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{queue: nil, size: k}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	} else {
		this.queue = append([]int{value}, this.queue...)
		return true
	}
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	} else {
		this.queue = append(this.queue, value)
		return true
	}
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	} else {
		this.queue = this.queue[1:]
		return true
	}
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	} else {
		this.queue = this.queue[:len(this.queue)-1]
		return true
	}
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.queue[0]
	}
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.queue[len(this.queue)-1]
	}
}

func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.queue) == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return len(this.queue) == this.size
}
func main() {
	circularDeque := Constructor(3) // 设置容量大小为3
	fmt.Println(circularDeque.InsertLast(1))
	fmt.Println(circularDeque.InsertLast(2))
	fmt.Println(circularDeque.InsertFront(3))
	fmt.Println(circularDeque.InsertFront(4))
	fmt.Println(circularDeque.GetRear())
	fmt.Println(circularDeque.IsFull())
	fmt.Println(circularDeque.DeleteLast())
	fmt.Println(circularDeque.InsertFront(4))
	fmt.Println(circularDeque.GetFront())
}
