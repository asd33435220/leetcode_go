package main

import "fmt"

func main() {

	res := modifyString("??yw?ipkj?")
	fmt.Println(res)
}

func modifyString(s string) string {
	byteArr := []byte(s)
	for i := 0; i < len(byteArr); i++ {
		lastChar := 96
		nextChar := 96
		if string(byteArr[i]) == "?" {
			if i > 0 {
				lastChar = int(byteArr[i-1])
			}
			if i < len(s)-1 {
				nextChar = int(byteArr[i+1])
			}
			char := 97
			for char == lastChar || char == nextChar {
				char++
			}
			byteArr[i] = uint8(char)
		} else {
			byteArr[i] = s[i]
		}
	}
	return string(byteArr)
}
