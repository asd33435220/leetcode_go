package main

import (
	"fmt"
	"math"
)

func main() {
	result := bulbSwitch(10)
	fmt.Println(result)
}

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n) + 0.5))
}
