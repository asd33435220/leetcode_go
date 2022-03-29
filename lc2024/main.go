package main

import "fmt"

func main() {
	key := "TFFT"
	res := maxConsecutiveAnswers(key, 1)
	fmt.Println("res", res)
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	var left, count, max, number int
	for i := 0; i < len(answerKey); i++ {
		c := answerKey[i]
		if c == 'T' {
			number++
		} else {
			number++
			count++
		}
		for count > k {
			if answerKey[left] == 'F' {
				count--
			}
			number--
			left++
		}
		if number > max {
			max = number
		}
	}
	left = 0
	count = 0
	number = 0
	for i := 0; i < len(answerKey); i++ {
		c := answerKey[i]
		if c == 'F' {
			number++
		} else {
			number++
			count++
		}

		for count > k {
			if answerKey[left] == 'T' {
				count--
			}
			number--
			left++
		}
		if number > max {
			max = number
		}
	}
	return max
}
