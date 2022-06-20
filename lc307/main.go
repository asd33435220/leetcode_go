package main

func main() {

}

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		arr: nums,
	}
}

func (this *NumArray) Update(index int, val int) {
	this.arr[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	result := 0
	for i := left; i <= right; i++ {
		result += this.arr[i]
	}
	return result
}
