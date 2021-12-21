package main

import "fmt"

func main() {
	trust := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	res := findJudge(4, trust)
	fmt.Println(res)
}

func findJudge(n int, trust [][]int) int {
	in := make([]int, n+1)
	out := make([]int, n+1)
	for i := 0; i < len(trust); i++ {
		out[trust[i][0]]++
		in[trust[i][1]]++
	}
	for i := 1; i < len(in); i++ {
		if in[i] == n-1 && out[i] == 0 {
			return i
		}
	}
	return -1
}
