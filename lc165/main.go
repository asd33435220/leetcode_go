package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	c1 := strings.Count(version1, ".")
	c2 := strings.Count(version2, ".")
	if c1 > c2 {
		for i := 0; i < c1-c2; i++ {
			version2 = version2 + ".0"
		}
	} else {
		for i := 0; i < c2-c1; i++ {
			version1 = version1 + ".0"
		}
	}
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	res1 := 0
	for i := 0; i < len(v1); i++ {
		number, _ := strconv.Atoi(v1[i])
		res1 += number * int(math.Pow(10, float64(len(v1)-1-i)))
	}
	res2 := 0
	for i := 0; i < len(v2); i++ {
		number, _ := strconv.Atoi(v2[i])
		res2 += number * int(math.Pow10(len(v2)-1-i))
	}
	if res1 > res2 {
		return 1
	} else if res1 < res2 {
		return -1
	} else {
		return 0
	}
}
func main() {
	res := compareVersion("0.1", "0.0.1")
	fmt.Println("res", res)
}
