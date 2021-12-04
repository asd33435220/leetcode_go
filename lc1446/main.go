package main

import "fmt"

func maxPower(s string) (ans int) {
	lastChar := s[0]
	max := 1
	ans = 1
	for i := 1; i < len(s); i++ {
		if s[i] == lastChar {
			max++
			if max > 1 && max > ans {
				ans = max
			}
		} else {
			max = 1
		}
		lastChar = s[i]
	}
	return
}

func main() {
	res := maxPower("cc")
	fmt.Println("Res", res)
}
