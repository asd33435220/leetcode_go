package main

import (
	"fmt"
	"strings"
)

func main() {
	res := defangIPaddr("1.1.1.1")
	fmt.Println(res)

}

func defangIPaddr(address string) string {
	arr := strings.Split(address, ".")
	return strings.Join(arr, "[.]")
}
