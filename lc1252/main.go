package main

import (
	"fmt"
)

func main() {
	indices := [][]int{{0, 1}, {1, 1}}
	res := oddCells(2, 3, indices)
	fmt.Println(res)
}
func oddCells(m int, n int, indices [][]int) (res int) {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}

	for _, indie := range indices {
		row, col := indie[0], indie[1]
		for rowKey, rows := range matrix {
			for colKey, _ := range rows {
				if row == rowKey {
					matrix[rowKey][colKey]++
				}
				if colKey == col {
					matrix[rowKey][colKey]++
				}
			}
		}
	}
	for _, rows := range matrix {
		for _, v := range rows {
			if v%2 == 1 {
				res++
			}
		}
	}
	return
}
