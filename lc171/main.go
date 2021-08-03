package main

import "fmt"

func main() {
	res := titleToNumber("AB")
	fmt.Println("res", res)
}
func titleToNumber(columnTitle string) int {
	byteArr := []byte(columnTitle)
	mul := 1
	res := 0
	for i := len(byteArr) - 1; i >= 0; i-- {
		count := int(byteArr[i]) - 64
		res += count * mul
		mul *= 26
	}
	return res
}
