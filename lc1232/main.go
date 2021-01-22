package main

import (
	"fmt"
	"math"
)

func main() {
	coordinates := [][]int{{2, 1}, {4, 2}, {6, 3}}
	res := checkStraightLine(coordinates)
	fmt.Println(res)
}
func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) <= 2 {
		return true
	}
	firstCoordinate := coordinates[0]
	secondCoordinate := coordinates[1]
	var k float32
	if firstCoordinate[0] - secondCoordinate[0] == 0{
		k =float32(firstCoordinate[0]-secondCoordinate[0]) / float32(firstCoordinate[1]-secondCoordinate[1])
		for i:=1;i<len(coordinates);i++ {
			if firstCoordinate[1]-coordinates[i][1] == 0 {
				return false
			}
			var thisK float32
			thisK = float32(firstCoordinate[0]-coordinates[i][0])/float32(firstCoordinate[1]-coordinates[i][1])
			if math.Abs(float64(thisK - k))>0.00001{
				return false
			}
		}
	}else{
		k =float32(firstCoordinate[1]-secondCoordinate[1]) / float32(firstCoordinate[0]-secondCoordinate[0])
		for i:=1;i<len(coordinates);i++ {
			if firstCoordinate[0]-coordinates[i][0] == 0 {
				return false
			}
			var thisK float32
			thisK = float32(firstCoordinate[1]-coordinates[i][1])/float32(firstCoordinate[0]-coordinates[i][0])
			if math.Abs(float64(thisK - k))>0.00001{
				return false
			}
		}
	}
	return true
}
