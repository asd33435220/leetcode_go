package main

import (
	"fmt"
	"math"
)

type MaxQueue struct {
	MaxValue int
	MaxIndex int
	Queue    []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		MaxValue: -1,
		MaxIndex: -1,
		Queue:    make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.Queue) == 0 {
		return -1
	}
	return this.MaxValue
}

func (this *MaxQueue) Push_back(value int) {
	if value >= this.MaxValue {
		this.MaxValue = value
		this.MaxIndex = len(this.Queue)
	}
	this.Queue = append(this.Queue, value)
}

func getMax(list []int) (index, v int) {
	max := math.MinInt
	for key, value := range list {
		if value > max {
			max = value
			index = key
		}
	}
	return index, max
}

func (this *MaxQueue) Pop_front() int {
	if len(this.Queue) == 0 {
		return -1
	} else {
		front := this.Queue[0]
		this.Queue = this.Queue[1:]
		if this.MaxIndex > 0 {
			this.MaxIndex--
		} else if this.MaxIndex == 0 {
			index, value := getMax(this.Queue)
			this.MaxIndex = index
			this.MaxValue = value
		}
		return front
	}
}

func main() {
	maxQueue := Constructor()
	maxQueue.Push_back(10)
	res := maxQueue.Max_value()
	fmt.Println("res", res)
	res = maxQueue.Pop_front()
	fmt.Println("res", res)
	maxQueue.Push_back(7)
	res = maxQueue.Max_value()
	fmt.Println("res", res)
	maxQueue.Push_back(99)
}
