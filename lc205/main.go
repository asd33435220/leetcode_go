package main

import "fmt"

func main() {
	s := "aa"
	t := "ab"
	res := isIsomorphic(s, t)
	fmt.Println(res)
}

func isIsomorphic(s string, t string) bool {
	s2t := make(map[int8]int8)
	t2s := make(map[int8]int8)

	for i := 0; i < len(s); i++ {
		a := int8(s[i])
		b := int8(t[i])
		v1,ok1 := s2t[a]
		v2,ok2 := t2s[b]
		if !ok1 && !ok2 {
			s2t[a] = b
			t2s[b] = a
		}else{
			if v1 != b || v2 != a {
				return false
			}
		}
	}
	return true
}
