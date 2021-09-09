package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rand7() int {
	rand.Seed(time.Now().UnixNano())
	return int(rand.Int31n(7) + 1)
}
func rand10() int {
	a := rand7()
	b := rand7()
	for true {
		num := (a-1)*7 + b // rand 49
		if num <= 40 {
			return num%10 + 1 // 拒绝采样
		}

		a = num - 40 // rand 9
		b = rand7()
		num = (a-1)*7 + b // rand 63
		if num <= 60 {
			return num%10 + 1
		}

		a = num - 60 // rand 3
		b = rand7()
		num = (a-1)*7 + b // rand 21
		if num <= 20 {
			return num%10 + 1
		}
	}
	return -1
}
func main() {
	for i := 0; i < 20; i++ {
		res := rand10()
		fmt.Println("res=", res)
	}

}
