package main

import (
	"fmt"
)

type MovingAverage struct {
	size  int
	queue []int
	sum   float64
}

func Constructor(size int) MovingAverage {
	return MovingAverage{size: size}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.queue) == this.size {
		this.sum -= float64(this.queue[0])
		this.queue = this.queue[1:]
	}
	this.sum += float64(val)
	this.queue = append(this.queue, val)
	return this.sum / float64(len(this.queue))
}

func main() {
	ma := Constructor(3)
	fmt.Println(ma.Next(1))
	fmt.Println(ma.Next(5))
	fmt.Println(ma.Next(10))
	fmt.Println(ma.Next(8))

}
