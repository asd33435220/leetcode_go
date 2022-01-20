package main

import "fmt"

type StoneType struct {
	type0 int
	type1 int
	type2 int
}

func stoneGameIX(stones []int) bool {
	s := &StoneType{
		type0: 0,
		type1: 0,
		type2: 0,
	}
	for _, value := range stones {
		if value%3 == 0 {
			s.type0++
		} else if value%3 == 1 {
			s.type1++
		} else {
			s.type2++
		}
	}
	if s.type0%2 == 0 {
		return s.type1 >= 1 && s.type2 >= 1
	} else {
		return s.type2-s.type1 > 2 || s.type1-s.type2 > 2
	}
}
func main() {
	chalk := []int{
		3, 4, 1, 2,
	}
	res := stoneGameIX(chalk)
	fmt.Println("res=", res)
}
