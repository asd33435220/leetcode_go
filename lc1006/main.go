package main

import "fmt"

func main()  {
	res := clumsy(10)
	fmt.Println(res)
}

func clumsy(N int) int {
	switch {
	case N == 1:
		return 1
	case N == 2:
		return 2
	case N == 3:
		return 6
	case N == 4:
		return 7

	case N%4 == 0:
		return N + 1
	case N%4 <= 2:
		return N + 2
	default:
		return N - 1
	}
}