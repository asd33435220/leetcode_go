package main

import (
	"container/heap"
	"fmt"
)

func main() {
	stock := Constructor()
	//["StockPrice","update","maximum","current","minimum","maximum","maximum","maximum","minimum","minimum","maximum","update","maximum","minimum","update","maximum","minimum","current","maximum","update","minimum","maximum","update","maximum","maximum","current","update","current","minimum","update","update","minimum","minimum","update","current","update","maximum","update","minimum"]
	//[[],[38,2308],[],[],[],[],[],[],[],[],[],[47,7876],[],[],[58,1866],[],[],[],[],[43,121],[],[],[40,5339],[],[],[],[32,5339],[],[],[43,6414]
	//,[49,9369],[],[],[36,3192],[],[48,1006],[],[53,8013],[]]
	stock.Update(38, 2308)
	stock.Maximum()
	stock.Current()
	stock.Maximum()
	stock.Maximum()
	stock.Maximum()
	stock.Maximum()

	stock.Update(47, 7876)
	stock.Update(58, 1866)
	stock.Update(43, 121)
	stock.Update(40, 5339)
	stock.Update(32, 5339)
	stock.Update(43, 6414)
	stock.Update(49, 3192)
	stock.Update(48, 1006)
	stock.Update(53, 8013)
	res := stock.Current()
	fmt.Println("res", res)
	res = stock.Maximum()
	fmt.Println("res", res)
	stock.Update(1, 3)
	res = stock.Maximum()
	fmt.Println("res", res)
	stock.Update(4, 2)
	res = stock.Minimum()
	fmt.Println("res", res)
}

type stockHeap [][]int

func (heap stockHeap) Len() int           { return len(heap) }
func (heap stockHeap) Top() []int         { return heap[0] }
func (heap stockHeap) Less(i, j int) bool { return heap[j][1] < heap[i][1] }
func (heap stockHeap) Swap(i, j int)      { heap[i], heap[j] = heap[j], heap[i] }
func (heap *stockHeap) Pop() interface{} {
	old := *heap
	n := len(old)
	x := old[n-1]
	*heap = old[0 : n-1]
	return x
}
func (heap *stockHeap) Push(stock interface{}) {
	*heap = append(*heap, stock.([]int))
}

type StockPrice struct {
	Map     map[int]int
	MaxHeap *stockHeap
	MinHeap *stockHeap
	currDay int
}

func Constructor() StockPrice {
	maxHeap := &stockHeap{}
	minHeap := &stockHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)
	return StockPrice{
		Map:     map[int]int{},
		MaxHeap: maxHeap,
		MinHeap: minHeap,
		currDay: 0,
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	this.Map[timestamp] = price
	heap.Push(this.MaxHeap, []int{timestamp, price})
	heap.Push(this.MinHeap, []int{timestamp, -price})
	if timestamp > this.currDay {
		this.currDay = timestamp
	}

}

func (this *StockPrice) Current() int {
	return this.Map[this.currDay]
}

func (this *StockPrice) Maximum() int {
	for true {
		max := this.MaxHeap.Top()
		maxDay := max[0]
		maxValue := max[1]
		if this.Map[maxDay] == maxValue {
			return maxValue
		}
		heap.Pop(this.MaxHeap)
	}
	return 0
}

func (this *StockPrice) Minimum() int {
	for true {
		min := this.MinHeap.Top()
		minDay := min[0]
		minValue := min[1]
		if this.Map[minDay] == -minValue {
			return -minValue
		}
		heap.Pop(this.MinHeap)
	}
	return 0
}
