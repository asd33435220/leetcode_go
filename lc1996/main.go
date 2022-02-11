package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{{7, 7}, {1, 2}, {9, 7}, {7, 3}, {3, 10}, {9, 8}, {8, 10}, {4, 3}, {1, 5}, {1, 5}}
	res := numberOfWeakCharacters(nums)
	fmt.Println("Res", res)

}

func numberOfWeakCharacters(properties [][]int) (ans int) {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][1] != properties[j][1] {
			return properties[i][1] > properties[j][1]
		} else {
			return properties[i][0] < properties[j][0]

		}
	})
	maxAtt := properties[0][0]
	for i := 1; i < len(properties); i++ {
		if properties[i][0] < maxAtt {
			ans++
		} else {
			maxAtt = properties[i][0]
		}
	}
	return
}
