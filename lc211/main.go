package main

import "fmt"

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func countDot(word string, index int) int {
	res := 0
	for i := index; i >= 0; i-- {
		if word[index] == '.' {
			res++
		}
	}
	return res
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
	for key, value := range word {
		if value == '.' {
			for i := 0; i < len(node.children); i++ {
				if node.children[i] == nil {
					continue
				}
				if word == "." && node.children[i].isEnd {
					return true
				} else {
					if node.children[i].Search(word[countDot(word, key):]) {
						return true
					}
				}
			}
			return false
		}
		value -= 'a'
		if node.children[value] == nil {
			return false
		}
		node = node.children[value]
	}
	return node.isEnd
}

type WordDictionary struct {
	tree Trie
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	this.tree.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.tree.Search(word)
}
func main() {
	wd := Constructor()
	wd.AddWord("at")
	wd.AddWord("an")
	wd.AddWord("add")
	wd.AddWord("and")

	res := wd.Search("a")
	fmt.Println("res1", res)
	res = wd.Search(".at")
	fmt.Println("res2", res)
	wd.AddWord("bat")
	res = wd.Search(".at")
	fmt.Println("res3", res)
	res = wd.Search("an.")
	fmt.Println("res4", res)
	res = wd.Search("a.d.")
	fmt.Println("res", res)
	res = wd.Search("b.")
	fmt.Println("res", res)
	res = wd.Search("a.d")
	fmt.Println("res", res)
	res = wd.Search(".")
	fmt.Println("res", res)
	//wd.AddWord("bad")
	//wd.AddWord("dad")
	//wd.AddWord("mad")
	//res := wd.Search("pad")
	//fmt.Println("res", res)
	//
	//res = wd.Search("bad")
	//fmt.Println("res", res)
	//
	//res = wd.Search(".ad")
	//fmt.Println("res", res)
	//
	//res = wd.Search("b..")
	//fmt.Println("res", res)

}
