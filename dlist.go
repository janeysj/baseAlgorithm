package main

import (
	"fmt"
	//"runtime"
	//"sync/atomic"
	//"unsafe"
)

// Doubly linked list

type ListNode struct {
	prenode  *ListNode
	nextnode *ListNode
	value    interface{}
}

type DList struct {
	head *ListNode
	tail *ListNode
}

func DListNew() *DList {
	l := new(DList)
	headnode := new(ListNode)
	tailnode := new(ListNode)
	headnode.nextnode = tailnode
	tailnode.prenode = headnode
	l.head = headnode
	l.tail = tailnode
	fmt.Printf("list is %v, tail is %v, head is %v", l, l.tail, l.head)
	return l
}

func (self *DList) InsertNode(val interface{}) error {
	fmt.Printf("===========InsertNode===========\n")
	newnode := ListNode{value: val}

	fmt.Printf("list tail is %v\n", self.tail)
	pre_tail := self.tail.prenode
	pre_tail.nextnode = &newnode
	self.tail.prenode = &newnode
	newnode.prenode = pre_tail
	newnode.nextnode = self.tail

	return nil
}

func (self *DList) Print() {
	fmt.Printf("===========DList Print===========\n")
	for p := self.head; p != nil; p = p.nextnode {
		if p.value != nil {
			fmt.Printf("%v  ", p.value)
		}
	}
	fmt.Printf("\n")
}

func (self *DList) ReversePrint() {
	fmt.Printf("===========DList Reverse Print===========\n")
	for p := self.tail; p != nil; p = p.prenode {
		if p.value != nil {
			fmt.Printf("%v  ", p.value)
		}
	}
	fmt.Printf("\n")
}

func main() {
	// q := QueueNew()
	// q.EnQueue("abc")
	q := DListNew()
	q.InsertNode("123")
	q.InsertNode("456")
	q.Print()
	q.ReversePrint()
}
