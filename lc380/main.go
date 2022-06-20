package main

import (
	"fmt"
	"math/rand"
)

func main() {
	set := Constructor()
	//["RandomizedSet","insert","insert","remove","insert","remove","getRandom"]
	//	[[],[0],[1],[0],[2],[1],[]]
	set.Insert(0)
	set.Insert(1)
	set.Insert(2)

	set.Insert(3)
	res := set.GetRandom()
	fmt.Println("res1", res)
	res = set.GetRandom()
	fmt.Println("res2", res)

}

type RandomizedSet struct {
	hash map[int]int
	nums []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hash: make(map[int]int),
		nums: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.hash[val]
	if ok {
		return false
	} else {
		this.hash[val] = len(this.nums)
		this.nums = append(this.nums, val)
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	key, ok := this.hash[val]
	if ok {
		last := len(this.nums) - 1
		lastVal := this.nums[last]
		if key != last {
			this.nums[key] = lastVal
			this.hash[lastVal] = key
		}
		delete(this.hash, val)
		this.nums = this.nums[:last]

		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
