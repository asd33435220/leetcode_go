package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) (res [][]int) {
	array := make([]int, 0)
	var dfs func(index, target int)
	dfs = func(index, target int) {
		if target < 0 {
			return
		}
		if target == 0 {
			// // 正确写法 生成新数组
			// newArr := make([]int, len(array))
			// copy(newArr, array)
			// res = append(res, newArr)

			// 匿名新数组 构建新数组指针
			fmt.Printf("1:%p\n", &array)
			tmp := append([]int{}, array...)
			fmt.Printf("2:%p\n", &tmp)
			res = append(res, tmp)

			// 错误写法 指针拷贝
			// newArr := array
			// res = append(res, newArr)
		}
		for i := index; i < len(candidates); i++ {
			array = append(array, candidates[i])
			dfs(i, target-candidates[i])
			array = array[:len(array)-1]
		}
	}
	dfs(0, target)
	return
}
func main() {
	candidates := []int{2, 3, 6, 7}
	res := combinationSum(candidates, 7)
	fmt.Println("res", res)
}
