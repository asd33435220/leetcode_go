package main

import "fmt"

func main() {
	A := []int{0}
	K := 34
	res := addToArrayForm(A, K)
	fmt.Println(res)
}
func addToArrayForm(A []int, K int) []int {
	var res []int
	n := len(A)
	for i := n - 1; i >= 0; i-- {
		sum := A[i] + K%10
		K /= 10
		if sum >= 10 {
			K++
			sum -= 10
		}
		res = append(res, sum)
	}
	for ; K > 0; K /= 10 {
		res = append(res, K%10)
	}
	SliceReverse(res)
	return res
}
func SliceReverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		tem1 := arr[i]
		tem2 := arr[len(arr)-i-1]
		arr[len(arr)-i-1] = tem1
		arr[i] = tem2
	}
}
