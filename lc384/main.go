package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slu := Constructor([]int{1, 2, 3})
	fmt.Println(slu.Shuffle())

	slu.Reset()
	slu.Shuffle()
}

type Solution struct {
	Data []int
}

func Constructor(nums []int) Solution {
	return Solution{Data: nums}
}

func (this *Solution) Reset() []int {
	return this.Data
}

func (this *Solution) Shuffle() []int {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, len(this.Data))
	copy(data, this.Data)
	var result []int
	for i := 0; i < len(this.Data); i++ {
		random := rand.Int31n(int32(len(data)))
		result = append(result, data[random])
		if int(random+1) == len(data) {
			data = data[:len(data)-1]
		} else {
			data = append(data[:random], data[random+1:]...)
		}
	}
	return result
}
