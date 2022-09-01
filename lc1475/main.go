package main

import "fmt"

type Good struct {
	price int
	index int
}

func finalPrices(prices []int) []int {
	n := len(prices)
	stack := make([]*Good, 0)
	res := make([]int, n)
	for i, price := range prices {
		for len(stack) > 0 && stack[len(stack)-1].price >= price {
			top := stack[len(stack)-1]
			res[top.index] = top.price - price
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, &Good{
			price: price,
			index: i,
		})
	}
	for _, v := range stack {
		res[v.index] = v.price
	}
	return res
}
func main() {
	target := []int{8, 4, 6, 2, 3}
	fmt.Println("res", finalPrices(target))
}
