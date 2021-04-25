package main

import "fmt"

func numDecodings(s string) int {
	var a, b, c int
	a = 1
	for i := 1; i <= len(s); i++ {
		if i == 1 {
			if s[i-1] != '0' {
				b = 1
			}
			continue
		}
		if s[i-1] != '0' {
			c += b
		}
		if s[i-2] != '0' {
			number := (s[i-2]-'0')*10 + (s[i-1] - '0')
			if number <= 26 {
				c += a
			}
		}
		a = b
		b = c
		c = 0
	}
	return b
}

func main() {
	s := "2222"
	res := numDecodings(s)
	fmt.Println("res", res)
}
