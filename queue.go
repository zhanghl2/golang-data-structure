package main

import "fmt"

// queue represents a queue that holds a slice
type queue struct {
	items []int
}

// Enqueue adds a value at the end
func (q *queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue removes a value at the front
// and returns the removed value
func (q *queue) Dequeue() int {
	toremove := q.items[0]
	q.items = q.items[1:]
	return toremove
}
func main() {
	myQueue := queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(1)
	myQueue.Enqueue(4)
	myQueue.Enqueue(3)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
