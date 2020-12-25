package main

import "fmt"

func main()  {
	s := "loveleetcode"
	res:=firstUniqChar(s)
	fmt.Println(res)
}

func firstUniqChar(s string) int {
	myMap := make(map[int32]int)
	for _,value := range s {
		v,ok := myMap[value]
		if !ok {
			myMap[value] = 1
		}else{
			myMap[value] = v+1
		}
	}
	for key,value := range s {
		_,ok := myMap[value]
		if !ok {
			myMap[value] = 1
		}else{
			if myMap[value] == 1 {
				return key
			}else{
				continue
			}
		}
	}
	return -1
}
