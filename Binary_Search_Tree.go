package main

import "fmt"

var count int

// Node represents the componets of a binary search tree
type Node struct {
	Key   int
	left  *Node
	right *Node
}

// Insert will add a node to the tree
// the key to add should not be already int the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.right == nil {
			n.right = &Node{Key: k}
		} else {
			n.right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.left == nil {
			n.left = &Node{Key: k}
		} else {
			n.left.Insert(k)
		}
	}
}

// Search will take in a key value
// and return true if there is a node with that value
func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}
	if n.Key < k {
		//move right
		return n.right.Search(k)
	} else if n.Key > k {
		//move left
		return n.left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree)
	tree.Insert(50)
	tree.Insert(200)

	fmt.Println(tree)
	fmt.Println(tree.Search(100))
	fmt.Println(count)
}
