package main

import (
	"fmt"
	"math"
)

func main() {
	res := poorPigs(1000, 15, 60)
	fmt.Println(res)
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(minutesToTest/minutesToDie+1))))
}
