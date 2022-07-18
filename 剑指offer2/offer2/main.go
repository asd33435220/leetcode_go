package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	x, _ := strconv.ParseInt(a, 2, 10)
	y, _ := strconv.ParseInt(b, 2, 10)
	return strconv.FormatInt(int64(x+y), 2)
}
func main() {
	res := addBinary("11", "10")

	fmt.Println(res)
}
