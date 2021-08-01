package maim

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}
	var max,now,left,right int
	right = 1
	for left <= len(nums){

	}
	return max
}
func main()  {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	res := maxSubArray(nums)
	fmt.Println("res",res)
}
