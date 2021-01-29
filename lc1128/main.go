package main

import (
	"fmt"
	"strconv"
)

func main() {
	dominoes := [][]int{{1, 1},
		{2, 2},
		{1, 1},
		{1, 2},
		{1, 2},
		{1, 1}}
	res := numEquivDominoPairs(dominoes)
	fmt.Println(res)

}

func numEquivDominoPairs(dominoes [][]int) int {
	myMap := make(map[string]int)
	res := 0
	I := func(x int) int {
		return x * (x - 1) / 2
	}
	for i := 0; i < len(dominoes); i++ {
		if dominoes[i][0] > dominoes[i][1] {
			dominoes[i][1], dominoes[i][0] = dominoes[i][0], dominoes[i][1]
		}
		str := strconv.Itoa(dominoes[i][0]) + strconv.Itoa(dominoes[i][1])
		value, ok := myMap[str]
		if !ok {
			myMap[str] = 1
		} else {
			myMap[str] = value + 1
		}
	}

	for _, value := range myMap {
		if value > 1 {
			res += I(value)
		}
	}
	return res
}
