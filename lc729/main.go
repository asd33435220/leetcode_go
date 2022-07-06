package main

import (
	"fmt"
)

type pair struct {
	start int
	end   int
}

type MyCalendar []pair

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, p := range *this {
		if p.start < end && p.end > start {
			return false
		}
	}
	*this = append(*this, pair{
		start,
		end,
	})
	return true
}
func main() {

	cal := Constructor()
	res := cal.Book(47, 50)
	fmt.Println("res", res)
	res = cal.Book(33, 41)
	fmt.Println("res", res)
	res = cal.Book(39, 45)
	fmt.Println("res", res)
	res = cal.Book(33, 42)
	fmt.Println("res", res)
	res = cal.Book(25, 32)
	fmt.Println("res", res)
	res = cal.Book(26, 35)
	fmt.Println("res", res)
}
