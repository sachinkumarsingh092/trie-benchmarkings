package main

import "fmt"

type Node struct {
	child  map[string]*Node
	stride int
	isEnd  bool
}

type Trie struct {
	root *Node
}

func createNode(stride int) *Node {
	return &Node{
		child:  make(map[string]*Node),
		stride: stride,
		isEnd:  false,
	}
}

func CreateTrie(stride int) *Trie {
	return &Trie{
		root: createNode(stride),
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	stride := t.root.stride

	for i := 0; i < len(word); i += stride {
		var key string
		if i+stride > len(word) {
			key = word[i:len(word)]
		} else {
			key = word[i : i+stride]
		}

		if node.child[key] == nil {
			newNode := createNode(stride)
			node.child[key] = newNode
		}
		node = node.child[key]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	stride := t.root.stride

	for i := 0; i < len(word); i += stride {
		var key string
		if i+stride > len(word) {
			key = word[i:len(word)]
		} else {
			key = word[i : i+stride]
		}

		if node.child[key] != nil {
			node = node.child[key]
			if node.isEnd {
				return true
			}
		}

	}

	return false
}

func main() {
	trie := CreateTrie(2)
	trie.Insert("lubna")
	trie.Insert("booboo")
	trie.Insert("luboo")
	trie.Insert("honeybee")
	//trie.Delete("luboo")
	fmt.Println(trie.Search("luboo"))
	//fmt.Println(trie.HasPrefix("lubu"))
	//fmt.Println(trie.HasPrefix("honey"))
}
