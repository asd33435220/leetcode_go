package main

import (
	"fmt"
	"sort"
)

type myStruct struct {
	power float32
	row int
}
func main() {
	res := kWeakestRows([][]int{{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}}, 3)
	fmt.Println("res",res)
}
func kWeakestRows(mat [][]int, k int) []int {
	msSlice := make([]myStruct,0)
	res := make([]int,0)

	for i := 0; i < len(mat); i++ {
		ms := myStruct{
			power: float32(getValue(mat[i]))+float32(i+1)/1000,
			row:   i,
		}
		msSlice = append(msSlice,ms)
	}
	sort.Slice(msSlice, func(i, j int) bool {
		return msSlice[i].power < msSlice[j].power
	})
	for i := 0; i < k; i++ {
		res = append(res,msSlice[i].row)
	}
	return res
}
func getValue(row []int) (result int) {
	for _, value := range row {
		if value == 1 {
			result++
		}else{
			break
		}
	}
	return result
}
