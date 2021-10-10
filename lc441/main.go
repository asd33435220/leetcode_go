package main

func arrangeCoins(n int) int {
	result := 0
	i := 1
	for n > 0 {
		n -= i
		i++
		if n >= 0 {
			result++
		}
	}
	return result
}
func main() {
	arrangeCoins(5)
}
