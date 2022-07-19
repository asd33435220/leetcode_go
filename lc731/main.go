package main

import "fmt"

type Pair struct {
	start int
	end   int
}
type MyCalendarTwo struct {
	booked     []Pair
	overBooked []Pair
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		booked:     nil,
		overBooked: nil,
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, pair := range this.overBooked {
		if !(pair.start >= end || pair.end <= start) {
			return false
		}
	}
	for _, pair := range this.booked {
		if !(pair.start >= end || pair.end <= start) {
			this.overBooked = append(this.overBooked, Pair{max(pair.start, start), min(pair.end, end)})
		}
	}
	this.booked = append(this.booked, Pair{start, end})
	return true
}
func main() {
	mc := Constructor()
	fmt.Println(mc.Book(10, 20))
	fmt.Println(mc.Book(50, 60))
	fmt.Println(mc.Book(10, 40))
	fmt.Println(mc.Book(5, 15))
	fmt.Println(mc.Book(5, 10))
	fmt.Println(mc.Book(25, 55))
}
