package main

import "fmt"

type word struct {
	A int
	E int
	I int
	O int
	U int
}

func add(number ...int) (res int) {
	for _, a := range number {
		res = (res + a) % (1e9 + 7)
	}
	return
}

func countVowelPermutation(n int) int {
	a := &word{
		A: 1,
		E: 1,
		I: 1,
		O: 1,
		U: 1,
	}
	for i := 1; i < n; i++ {
		b := &word{
			A: add(a.E),
			E: add(a.A, a.I),
			I: add(a.A, a.E, a.O, a.U),
			O: add(a.I, a.U),
			U: add(a.A),
		}
		a = b
	}
	return add(a.A, a.E, a.I, a.O, a.U)
}
func main() {
	res := countVowelPermutation(144)
	fmt.Println("res", res)
}
