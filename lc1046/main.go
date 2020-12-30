package main

import (
	"fmt"
	"sort"
)

type Heap struct {
	sort.IntSlice
}

func (this *Heap) GetSize() int {
	return this.IntSlice.Len()
}
func (this *Heap) GetTop() int {
	return this.IntSlice[0]
}
func (this *Heap) Pop() int {
	this.Swap(0, len(this.IntSlice)-1)
	top := this.IntSlice[len(this.IntSlice)-1]
	this.IntSlice = this.IntSlice[:len(this.IntSlice)-1]
	this.SiftDown(0)
	return top
}
func (this *Heap) Push(element int) {
	this.IntSlice = append(this.IntSlice, element)
	index := this.IntSlice.Len() - 1
	this.SiftUp(index)
}
func (this *Heap) SiftUp(index int) {
	for index != 0 {
		if this.IntSlice[(index-1)/2] > this.IntSlice[index] {
			break
		}
		this.Swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}
func (this *Heap) SiftDown(index int) {
	leftChild := (index+1)*2 - 1
	rightChild := (index + 1) * 2
	for index != this.IntSlice.Len()-1 {
		if leftChild > this.IntSlice.Len()-1 {
			break
		}
		if rightChild > this.IntSlice.Len()-1 {
			if this.IntSlice[index] < this.IntSlice[leftChild] {
				this.Swap(index, leftChild)
				index = leftChild
				continue
			} else {
				break
			}
		}
		if this.IntSlice[index] > this.IntSlice[leftChild] && this.IntSlice[index] > this.IntSlice[rightChild] {
			break
		}
		if this.IntSlice[leftChild] > this.IntSlice[rightChild] {
			this.Swap(index, leftChild)
			index = leftChild
		} else {
			this.Swap(index, rightChild)
			index = rightChild
		}
		leftChild = (index+1)*2 - 1
		rightChild = (index + 1) * 2
	}
}
func (this *Heap) Less(i, j int) bool {
	return this.IntSlice.Less(i, j)
}

func main() {
	arr := []int{10, 4, 2, 10}
	res := lastStoneWeight(arr)
	fmt.Println(res)
}
func lastStoneWeight(stones []int) int {
	myHeap := &Heap{[]int{}}
	len := len(stones)
	for i := 0; i < len; i++ {
		top := stones[0]
		stones = stones[1:]
		myHeap.Push(top)
	}
	for myHeap.IntSlice.Len() > 1 {
		top1 := myHeap.Pop()
		top2 := myHeap.Pop()
		if top1-top2 == 0 {
			continue
		} else {
			myHeap.Push(top1 - top2)
		}
	}
	if myHeap.IntSlice.Len() == 1 {
		return myHeap.IntSlice[0]
	} else if myHeap.IntSlice.Len() == 0 {
		return 0
	} else {
		return 0
	}
}
