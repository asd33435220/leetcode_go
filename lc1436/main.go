package main

import "fmt"

func destCity(paths [][]string) string {
	myMap := make(map[string]bool)
	for i := 0; i < len(paths); i++ {
		myMap[paths[i][0]] = true
	}
	for i := 0; i < len(paths); i++ {
		if !myMap[paths[i][1]] {
			return paths[i][1]
		}
	}
	return ""
}

func main() {
	paths := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	res := destCity(paths)
	fmt.Println("Res", res)
}
