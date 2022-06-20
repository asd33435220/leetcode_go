package main

import "fmt"

func setResult(resource chan<- int, number int) {
	for i := 1; i < number; i++ {
		resource <- i
	}
	close(resource)
}
func count(resource, result chan int, exit chan bool) {
	ok := true
	for ok {
		var count int
		flag := true
		count, ok = <-resource
		for i := 2; i < count; i++ {
			if count%i == 0 {
				flag = false
				break
			}
		}
		if flag && count > 2 {
			result <- count
		}
	}
	fmt.Println("携程退出")
	exit <- true
}
func main() {
	source := make(chan int, 100000)
	result := make(chan int, 10000)
	exit := make(chan bool, 16)
	go setResult(source, 100000)
	for i := 0; i < 16; i++ {
		go count(source, result, exit)
	}
	go func() {
		for i := 0; i < 16; i++ {
			<-exit
		}
		close(result)
	}()
	for value := range result {
		fmt.Println("质数是", value)
	}
}
