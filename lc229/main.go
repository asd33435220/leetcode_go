package main

func majorityElement(nums []int) (res []int) {
	countMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		countMap[nums[i]]++
	}
	for key, value := range countMap {
		if value > len(nums)/3 {
			res = append(res, key)
		}
	}
	return res
}
func main() {

}
