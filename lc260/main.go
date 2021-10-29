package main

func singleNumber(nums []int) (res []int) {
	myMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		myMap[nums[i]]++
	}

	for key, value := range myMap {
		if value == 1 {
			res = append(res, key)
		}
	}
	return

}
func main() {
	singleNumber([]int{4, 1, 2, 3})
}
