package main

import (
	"fmt"
	"strconv"
)

func binaryGap(n int) int {
	binaryStr := strconv.FormatInt(int64(n), 2)
	lastIndex := -1
	maxGap := 0
	for i := 0; i < len(binaryStr); i++ {
		char := binaryStr[i]
		if char == '1' {
			if gap := i - lastIndex; lastIndex >= 0 && gap > maxGap {
				maxGap = gap
			}
			lastIndex = i
		}
	}
	return maxGap
}
func main() {

	res := binaryGap(5)
	fmt.Println("Res", res)
}
