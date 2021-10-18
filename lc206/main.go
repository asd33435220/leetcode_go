package main

import "fmt"

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, value := range word {
		value -= 'a'
		if node.children[value] == nil {
			node.children[value] = &Trie{}
		}
		node = node.children[value]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, value := range word {
		value -= 'a'
		if node.children[value] == nil {
			return false
		}
		node = node.children[value]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, value := range prefix {
		value -= 'a'
		if node.children[value] == nil {
			return false
		}
		node = node.children[value]
	}
	return true
}
func main() {
	tree := Constructor()
	tree.Insert("apple")
	res := tree.Search("apple")
	fmt.Println("res", res)

	res = tree.Search("app")
	fmt.Println("res", res)

	res = tree.StartsWith("app")
	fmt.Println("res", res)

	tree.Insert("app")
	res = tree.Search("app")
	fmt.Println("res", res)

}
