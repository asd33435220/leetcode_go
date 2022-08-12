package main

import (
	"fmt"
)

type HashSlice struct {
	slice    []Slice
	sliceMap map[int]int
}
type Slice struct {
	arr []int
	max int
}

func groupThePeople(groupSizes []int) (res [][]int) {
	//var hashSlice HashSlice
	//hashSlice.sliceMap = make(map[int]int)
	//hash := make(map[int]int)
	//for _, v := range groupSizes {
	//	hash[v]++
	//}
	//for k, v := range hash {
	//	size := v / k
	//	for i := 0; i < size; i++ {
	//		hashSlice.slice = append(hashSlice.slice, Slice{
	//			arr: nil,
	//			max: k,
	//		})
	//
	//		if _, ok := hashSlice.sliceMap[k]; !ok {
	//			hashSlice.sliceMap[k] = len(hashSlice.slice) - 1
	//		}
	//	}
	//}
	//for k, v := range groupSizes {
	//	index := hashSlice.sliceMap[v]
	//	hashSlice.slice[index].arr = append(hashSlice.slice[index].arr, k)
	//	if len(hashSlice.slice[index].arr) == hashSlice.slice[index].max {
	//		hashSlice.sliceMap[v]++
	//	}
	//}
	//for i := 0; i < len(hashSlice.slice); i++ {
	//	res = append(res, hashSlice.slice[i].arr)
	//}
	//return
	hash := make(map[int][]int)
	for k, v := range groupSizes {
		hash[v] = append(hash[v], k)
	}
	for k, v := range hash {
		for i := 0; i < len(v); i += k {
			res = append(res, v[i:i+k])
		}
	}
	return
}

func main() {
	groupSize := []int{3, 3, 3, 3, 3, 1, 3}
	res := groupThePeople(groupSize)
	fmt.Println(res)
}
