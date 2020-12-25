package main

import "fmt"

func main() {

	res := generate(5)
	fmt.Println(res)
}
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		if i == 0 {
			result[i] = append(result[i], 1)
		}
		if i == 1 {
			result[i] = append(result[i], 1, 1)
		}
		if i > 1 {
			for j := 0; j <= i; j++ {
				if j == 0 || j == i {
					result[i] = append(result[i], 1)
				} else {
					result[i] = append(result[i], result[i-1][j]+result[i-1][j-1])
				}
			}
		}
	}
	return result
}
