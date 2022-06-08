package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	x      float64
	y      float64
	radius float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		radius: radius,
		x:      x_center,
		y:      y_center,
	}
}

func (this *Solution) RandPoint() []float64 {
	for {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1
		if x*x+y*y < 1 {
			return []float64{x*this.radius + this.x, y*this.radius + this.y}
		}
	}

}

func main() {
	C := Constructor(1.1, 2.2, 3.3)
	res := C.RandPoint()
	fmt.Println("res", res)
}
