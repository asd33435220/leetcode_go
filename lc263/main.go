package main

func isUgly(n int) bool {
	if n == 0 {
		return false
	}
	arr := [3]int{2, 3, 5}
	for _, value := range arr {
		for n%value == 0 {
			n /= value
		}
	}
	return n == 1
}
func main() {
	isUgly(12)
}
