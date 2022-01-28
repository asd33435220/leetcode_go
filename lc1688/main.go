package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	res := numberOfMatches(7)
	fmt.Println("res", res)
}

func numberOfMatches(n int) (res int) {
	for n > 1 {
		if n%2 == 0 {
			res += n / 2
			n = n / 2
		} else {
			res += (n - 1) / 2
			n = (n-1)/2 + 1
		}
	}
	return
}
