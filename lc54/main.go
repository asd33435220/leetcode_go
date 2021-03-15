package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}}
	res := spiralOrder(arr)
	fmt.Println(res)
}

func spiralOrder(matrix [][]int) []int {
	var res []int
	var walk func([][]int, int)
	walk = func(matrix [][]int, k int) {
		if len(matrix)-2*k == 0 || len(matrix[0])-2*k == 0 {
			return
		} else if len(matrix)-2*k == 1 {
			for i := k; i < len(matrix[0])-k; i++ {
				res = append(res, matrix[k][i])
			}
			return
		} else if len(matrix[0])-2*k == 1 {
			for i := k; i < len(matrix)-k; i++ {
				res = append(res, matrix[i][k])
			}
			return
		}
		for i := k; i < len(matrix[0])-k; i++ {
			res = append(res, matrix[k][i])
		}
		for i := k + 1; i < len(matrix)-k; i++ {
			res = append(res, matrix[i][len(matrix[0])-k-1])
		}
		for i := len(matrix[0]) - k - 2; i > k-1; i-- {
			res = append(res, matrix[len(matrix)-k-1][i])
		}
		for i := len(matrix) - k - 2; i > k; i-- {
			res = append(res, matrix[i][k])
		}
		walk(matrix, k+1)
	}
	walk(matrix,0)
	return res
}
