package main

import "fmt"

type Node struct {
	child map[rune]*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

func createNode() *Node {
	return &Node{
		child: make(map[rune]*Node),
		isEnd: false,
	}
}

func CreateTrie() *Trie {
	return &Trie{
		root: createNode(),
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if node.child[char] == nil {
			newNode := createNode()
			node.child[char] = newNode
		}
		node = node.child[char]
	}
	node.isEnd = true
}

func (t *Trie) Delete(word string) {
	node := t.root
	delSlice := []*Node{node}

	for _, char := range word {
		if node.child[char] != nil {
			node = node.child[char]
			delSlice = append(delSlice, node)

			if node.isEnd {
				node.isEnd = false
			}
		}
	}

	for i := len(delSlice) - 1; i >= 0; i-- {
		current := delSlice[i]
		parent := delSlice[i-1]
		if len(current.child) == 0 && current.isEnd == false {
			delete(parent.child, rune(word[i-1]))
		} else {
			break
		}
	}
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if node.child[char] != nil {
			node = node.child[char]
			if node.isEnd {
				return true
			}

		}
	}
	return false
}

func (t *Trie) HasPrefix(word string) bool {
	node := t.root
	for _, char := range word {
		if node.child[char] != nil {
			node = node.child[char]
		} else {
			return false
		}
	}
	return true
}

func main() {
	trie := CreateTrie()
	trie.Insert("lubna")
	trie.Insert("booboo")
	trie.Insert("luboo")
	trie.Insert("honeybee")
	// trie.Delete("luboo")
	fmt.Println(trie.Search("luboo"))
	fmt.Println(trie.HasPrefix("lubu"))
	fmt.Println(trie.HasPrefix("honey"))
}
