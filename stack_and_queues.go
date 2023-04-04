// LIFO:stack--last in first out -->books
// FIFO:queue --first in first out -->queue
package main

import "fmt"

// stack represents stack that hold a slice
type Stack struct {
	items []int
}

// push will add a value at the end
func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

// pop will remove a value at the end
// and return the removed value
func (s *Stack) pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	mystack := Stack{}
	fmt.Println(mystack)
	mystack.push(1)
	mystack.push(2)
	mystack.push(9)
	fmt.Println(mystack)
	mystack.pop()
	fmt.Println(mystack)
}
