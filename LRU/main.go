package main

import "fmt"

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
type Node struct {
	key  int
	val  int
	next *Node
	prev *Node
}
type DoubleList struct {
	head *Node
	tail *Node
	size int
}

func (this *DoubleList) addFirst(node *Node) {
	this.size++
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}
func (this *DoubleList) removeNode(node *Node) {
	this.size--
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (this *DoubleList) removeLast() *Node {
	this.size--
	node := this.tail.prev
	this.tail.prev.prev.next = this.tail
	this.tail.prev = this.tail.prev.prev
	return node
}
func (this *DoubleList) getSize() int {
	return this.size
}

type LRUCache struct {
	cap    int
	list   *DoubleList
	lruMap map[int]*Node
}

func LRU(operators [][]int, k int) []int {
	res := make([]int, 0)
	// write code here
	myLRU := &LRUCache{
		cap: k,
		list: &DoubleList{
			head: &Node{
				key: -1,
				val: -1,
			},
			tail: &Node{
				key: -1,
				val: -1,
			},
			size: 0,
		},
		lruMap: make(map[int]*Node),
	}
	myLRU.list.head.next = myLRU.list.tail
	myLRU.list.tail.prev = myLRU.list.head

	for i := 0; i < len(operators); i++ {
		if operators[i][0] == 1 {
			myLRU.Set(operators[i][1], operators[i][2])
		} else if operators[i][0] == 2 {
			res = append(res, myLRU.Get(operators[i][1]))
		}
	}
	return res
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.lruMap[key]; !ok {
		return -1
	} else {
		value := node.val
		this.Set(key, value)
		return value
	}
}
func (this *LRUCache) Set(key, val int) {
	newNode := &Node{
		key: key,
		val: val,
	}
	if node, ok := this.lruMap[key]; ok {
		this.list.removeNode(node)
		this.list.addFirst(newNode)
		this.lruMap[key] = newNode
	} else {
		if this.cap == this.list.getSize() {
			lastNode := this.list.removeLast()
			delete(this.lruMap, lastNode.key)
		}

		this.list.addFirst(newNode)
		this.lruMap[key] = newNode

	}
}
func main() {
	operators := [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}
	res := LRU(operators, 3)
	fmt.Println("res", res)
}
