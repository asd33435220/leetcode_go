package main

import "fmt"

func main() {
	res := 3
	fmt.Println(res)
}

type MyHashSet struct {
	MyMap map[int]int
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
	myMap := MyHashSet{
		MyMap:make(map[int]int),
	}
	return myMap
}


func (this *MyHashSet) Add(key int)  {
	this.MyMap[key] = 1
}


func (this *MyHashSet) Remove(key int)  {
	delete(this.MyMap,key)
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	_,ok := myMap[key]
	return ok
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */