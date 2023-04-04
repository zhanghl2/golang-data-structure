package main

import "fmt"

// AlphabetSize is the number of characters in the trie
const AlphabetSize = 26

// Node represents each node in the trie
type Node struct {
	children [26]*Node
	isEnd    bool
}

// Trie represents a trie and has a pointer
type Trie struct {
	root *Node
}

// InitTrie will create a new trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}

		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take in a word and return true is that word is included in the trie
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	myTrie := InitTrie()
	toAdd := []string{
		"ssss",
		"aaa",
		"aac",
	}
	for _, word := range toAdd {
		myTrie.Insert(word)
	}
	myTrie.Insert("zhangguoxiu")
	fmt.Println(myTrie.Search("zhangguoxiu"))

}
