package main

import "fmt"

func main() {

	res := countPrimes(1500000)
	fmt.Println(res)
}
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	if n == 3 {
		return 1
	}
	if n <= 5 {
		return 2
	}
	var count = 0
	var result = &count
	numChan := make(chan int, 1000)
	priChan := make(chan int, 120000)
	resultChan := make(chan bool, 16)
	go putNumbers(n, numChan)
	for i := 0; i < 16; i++ {
		go getIsPrime(numChan, priChan, resultChan, result)
	}
	for i := 0; i < 16; i++ {
		<-resultChan
	}

	return (*result) + 2
}
func putNumbers(k int, numChan chan int) {
	for i := 1; i < k; i++ {
		numChan <- i
	}
	close(numChan)
}
func getIsPrime(numChan chan int, priChan chan int, result chan bool, res *int) {
	for value := range numChan {
		var isBreak = false
		var isInLoop = false
		for i := 2; i*i <= value; i++ {
			isInLoop = true
			if value%i == 0 {
				isBreak = true
				break
			}
		}
		if !isBreak && isInLoop {
			priChan <- value
			(*res)++
		}
	}
	result <- true
}
