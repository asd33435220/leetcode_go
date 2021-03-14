package main

import "fmt"

func main() {
	res := 3
	fmt.Println(res)
}

type MyHashMap struct {
	MyMap map[int]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	myHash := MyHashMap{MyMap: make(map[int]int)}
	return myHash
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.MyMap[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	value, ok := this.MyMap[key]
	if ok {
		return value
	} else {
		return -1
	}
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	delete(this.MyMap, key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
