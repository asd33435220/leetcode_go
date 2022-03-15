package main

import "fmt"

func main() {
	s := "AB"
	res := convert(s, 1)
	fmt.Println(res)
}

func convert(s string, numRows int) (ans string) {
	if numRows == 1 {
		return s
	}
	n := len(s)
	matrix := make([][]byte, numRows)
	isDown := false
	currRow := 0
	for i := 0; i < n; i++ {
		if currRow == 0 || currRow == numRows-1 {
			isDown = !isDown
		}
		c := s[i]
		matrix[currRow] = append(matrix[currRow], c)
		if isDown {
			currRow++
		} else {
			currRow--
		}

	}
	for _, byteArr := range matrix {
		ans += string(byteArr)
	}
	return
}
