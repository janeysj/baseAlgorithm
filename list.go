package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"unsafe"
)

type ListNode struct {
	nextnode unsafe.Pointer
	value    interface{}
}

type List struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func ListNew() (*List) {
	l := new(List)
	tailnode := (unsafe.Pointer)(&ListNode{nextnode: nil, value: nil})
	headnode := (unsafe.Pointer)(&ListNode{nextnode: tailnode, value: nil})
	l.head = headnode
	l.tail = tailnode
	fmt.Printf("list is %v, tail is %v, head is %v", l, l.tail, l.head)
	return l
}

// The safe insert node operation that use CAS.
func (self *List) InsertNode(val interface{}) error {
	fmt.Printf("===========InsertNode===========\n")
	newnode := (unsafe.Pointer)(&ListNode{value: val})
	tail := self.tail
	//(*ListNode)(self.tail).nextnode = newnode
	//self.tail = newnode
	for {
		nextnode := (*ListNode)(self.tail).nextnode
		if tail != self.tail {
			runtime.Gosched()
			continue
		}
		if nextnode != nil {
			atomic.CompareAndSwapPointer(&self.tail, tail, nextnode)
			continue
		}
		if atomic.CompareAndSwapPointer(&(*ListNode)(self.tail).nextnode, nil, newnode) {
			break
		}
		runtime.Gosched()
	}
	atomic.CompareAndSwapPointer(&self.tail, tail, newnode)

	return nil
}

func (self *List) DelNode(val interface{}) error {
	fmt.Printf("===========Delete Node===========\n")
	p := self.head
	q := p
	for ; p != nil; p = (*ListNode)(p).nextnode {
		if (*ListNode)(p).value != nil {
			if (*ListNode)(p).value == val {
				// Found the node.
				break
			}
		}
		q = p
	}
	if p == nil {
		fmt.Printf("There no element which value is %v\n", val)
		return fmt.Errorf("There no element which value is %v", val)
	}
	(*ListNode)(q).nextnode = (*ListNode)(p).nextnode

	return nil
}

func (self *List) Print() {
	fmt.Printf("===========List Print===========\n")
	for p := self.head; p != nil; p = (*ListNode)(p).nextnode {
		if (*ListNode)(p).value != nil {
			fmt.Printf("%v  ", (*ListNode)(p).value)
		}
	}
	fmt.Printf("\n")
}

func main() {
	// q := QueueNew()
	// q.EnQueue("abc")
	q := ListNew()
	q.InsertNode("123")
	q.InsertNode("456")
	q.InsertNode("789")
	q.Print()

	q.DelNode("abc")

	q.Print()
}
