package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	n := len(timePoints)
	timeList := make([]int, n)
	for index, timePoint := range timePoints {
		timeArr := strings.Split(timePoint, ":")
		hour, _ := strconv.Atoi(timeArr[0])
		min, _ := strconv.Atoi(timeArr[1])
		timeList[index] += hour*60 + min
	}
	sort.Ints(timeList)
	gap := timeList[1] - timeList[0]
	max := gap
	for i := 2; i < n; i++ {
		if timeList[i]-timeList[i-1] < gap {
			gap = timeList[i] - timeList[i-1]
		}
		if timeList[i]-timeList[i-1] > max {
			max = timeList[i] - timeList[i-1]
		}
	}
	if timeList[n-1]-timeList[0] < gap {
		gap = timeList[n-1] - timeList[0]
	}
	if timeList[n-1]-timeList[0] > max {
		max = timeList[n-1] - timeList[0]
	}
	if 1440-max < gap {
		return 1440 - max
	}
	return gap
}

func main() {
	timePoints := []string{"00:00", "04:00", "22:00"}
	res := findMinDifference(timePoints)
	fmt.Println("res", res)
}
