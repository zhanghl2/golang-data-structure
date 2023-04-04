package main

import "fmt"

// ArraySize is the size of the hash table array
const ArraySize = 7

// HashTable structure will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket structure is a linked list in each slot of the hash table
type bucket struct {
	head *bucketNode
}

// bucket node structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take in a key and return true if that key stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will delete a key from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

//insert will take in a key ,create a node with the key and Insert the node in the bucket

func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("already exists")
	}
}

// search will take in a key and return true if the bucket has that key
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			//delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next

	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	list := []string{
		"ww",
		"qqq",
		"wwe",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}
	hashTable.Delete("ww")
	fmt.Println("ww", hashTable.Search("ww"))
	fmt.Println("qqq", hashTable.Search("qqq"))
}
