package main

import "fmt"

func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for i := 0; i < len(chalk); i++ {
		sum += chalk[i]
	}
	k = k % sum
	number := 0
	for k >= 0 {
		k -= chalk[number]
		number = (number + 1) % len(chalk)
	}
	if number == 0 {
		return len(chalk)
	}
	return (number - 1) % len(chalk)
}
func main() {
	chalk := []int{
		3, 4, 1, 2,
	}
	res := chalkReplacer(chalk, 25)
	fmt.Println("res=", res)
}
