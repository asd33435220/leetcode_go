package main

import "fmt"

func removeDuplicates(S string) string {
	var sli []uint8
	sli = append(sli, S[0])
	for i := 1; i < len(S); i++ {
		if len(sli) > 0 && sli[len(sli)-1] == S[i] {
			sli = sli[:len(sli)-1]
		} else {
			sli = append(sli, S[i])
		}
	}
	res := string(sli)
	return res
}

func main() {
	s := "abbaca"
	res := removeDuplicates(s)
	fmt.Print(res)
}
