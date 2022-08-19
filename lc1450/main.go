package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	res := 0
	for i, s := range startTime {
		e := endTime[i]
		if s <= queryTime && e >= queryTime {
			res++
		}
	}
	return res
}

func main() {
}
