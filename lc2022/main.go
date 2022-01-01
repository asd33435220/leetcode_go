package main

import "fmt"

func construct2DArray(original []int, m int, n int) (res [][]int) {
	if len(original) != m*n {
		return
	}
	count := 0
	for i := 0; i < m; i++ {
		row := make([]int, n)
		for j := 0; j < n; j++ {
			row[j] = original[count]
			count++
		}
		res = append(res, row)
	}
	return
}
func main() {
	chalk := []int{
		3, 4, 1, 2,
	}
	res := construct2DArray(chalk, 2, 2)
	fmt.Println("res=", res)
}
