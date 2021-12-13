package main

import (
	"fmt"
	"strings"
)

func main() {
	res := toLowerCase("TEST")
	fmt.Println(res)
}

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
