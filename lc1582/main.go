package main

import "fmt"

func numSpecial(mat [][]int) (res int) {
	m := len(mat)
	n := len(mat[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				sum := 0
				for k := 0; k < m; k++ {
					if mat[k][j] == 1 {
						sum++
					}
				}
				for k := 0; k < n; k++ {
					if mat[i][k] == 1 {
						sum++
					}
				}
				if sum == 2 {
					res++
				}
			}
		}
	}
	return
}

func main() {
	nums := [][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}
	res := numSpecial(nums)
	fmt.Println("res", res)
}
