package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := dayOfYear("2008-10-10")
	fmt.Println(res)
}

func dayOfYear(date string) (ans int) {
	dateArr := strings.Split(date, "-")
	yearF, _ := strconv.Atoi(dateArr[0])
	monthF, _ := strconv.Atoi(dateArr[1])
	dayF, _ := strconv.Atoi(dateArr[2])
	year := int(yearF)
	month := int(monthF)
	day := int(dayF)
	isLeapYear := (year%4 == 0 && year%100 != 0) || year%400 == 0
	for month > 1 {
		month--
		if month == 2 {
			if isLeapYear {
				ans += 29
			} else {
				ans += 28
			}
		} else if month < 8 {
			if month%2 == 0 {
				ans += 30
			} else {
				ans += 31
			}
		} else if month%2 == 0 {
			ans += 31
		} else {
			ans += 30
		}
	}
	ans += day
	return
}
