package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Log struct {
	fn      int
	time    int
	logType int // 0 start 1 end
}

func parseLog(log string) Log {
	arr := strings.Split(log, ":")
	fn, _ := strconv.Atoi(arr[0])
	time, _ := strconv.Atoi(arr[2])
	logType := 1
	if arr[1] == "start" {
		logType = 0
	}
	return Log{fn, time, logType}
}
func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	stack := make([]Log, 0)
	for i := 0; i < len(logs); i++ {
		myLog := parseLog(logs[i])
		if myLog.logType == 1 && len(stack) != 0 && stack[len(stack)-1].fn == myLog.fn {
			top := stack[len(stack)-1]
			gap := myLog.time - top.time + 1
			res[top.fn] += gap
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				continue
			}
			topIndex := stack[len(stack)-1].fn
			res[topIndex] -= gap
		} else {
			stack = append(stack, myLog)
		}
	}
	return res

}
func main() {
	logs := []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"}
	res := exclusiveTime(2, logs)
	fmt.Println(res)
}
