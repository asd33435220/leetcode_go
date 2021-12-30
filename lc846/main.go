package main

import (
	"fmt"
	"sort"
)

func main() {
	hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	res := isNStraightHand(hand, 3)
	fmt.Println(res)
}
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	myMap := make(map[int]int)
	for i := 0; i < len(hand); i++ {
		myMap[hand[i]]++
	}
	for i := 0; i < len(hand); i++ {
		if myMap[hand[i]] == 0 {
			continue
		}
		myMap[hand[i]]--
		for j := 1; j < groupSize; j++ {
			if myMap[hand[i]+j] < 1 {
				return false
			} else {
				myMap[hand[i]+j]--
			}
		}
	}
	return true
}
