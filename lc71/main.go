package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	stack := make([]string, 0)
	stack = append(stack, "root")
	for i := 0; i < len(arr); i++ {
		if arr[i] == "" || arr[i] == "." || arr[i] == "/" {
		} else if arr[i] == ".." {
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, arr[i])
		}
	}
	if len(stack) == 1 {
		return "/"
	} else {
		return "/" + strings.Join(stack[1:], "/")
	}
}
func main() {
	path := "/a/./b/../../c/"
	res := simplifyPath(path)
	fmt.Println("res", res)
}
