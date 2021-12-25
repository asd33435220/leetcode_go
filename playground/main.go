package main

import (
	"fmt"
	"unsafe"
)

type B struct {
	a int
	b int
}

type A struct {
	b1 *B
	b2 *B
}

func main() {
	var a A
	fmt.Println(a.b1)
	fmt.Println(unsafe.Sizeof(a))
}
