package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkLIst struct {
	head   *node
	length int
}

// 头插法
func (l *linkLIst) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}
func (l linkLIst) printListdata() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}
func (l *linkLIst) deletewithvalue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	pretoDelete := l.head
	for pretoDelete.next.data != value {
		if pretoDelete.next.next == nil {
			return
		}
		pretoDelete = pretoDelete.next
	}
	pretoDelete.next = pretoDelete.next.next
	l.length--
}

func main() {
	mylist := linkLIst{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	mylist.prepend(node1)
	mylist.prepend(node2)
	fmt.Println(mylist)
	mylist.printListdata()
}
