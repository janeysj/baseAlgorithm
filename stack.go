package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"unsafe"
)

// lock free stack
type Stack struct {
	head  unsafe.Pointer
	tail  unsafe.Pointer
	Reset func(interface{})
	New   func() interface{}
}

// one node in stack
type Node struct {
	val  interface{}
	next unsafe.Pointer
}

func StackNew() *Stack {
	stack := new(Stack)
	stack.head = unsafe.Pointer(new(Node))
	stack.tail = stack.head
	return stack
}

// stack functions
func (self *Stack) EnStack(val interface{}) {

	if self.Reset != nil {
		self.Reset(val)
	}
	newNode := unsafe.Pointer(&Node{val: val, next: nil})
	var tail, next unsafe.Pointer
	for {
		tail = self.tail
		next = ((*Node)(tail)).next
		if tail != self.tail {
			runtime.Gosched()
			continue
		}
		if next != nil {
			atomic.CompareAndSwapPointer(&(self.tail), tail, next)
			continue
		}
		if atomic.CompareAndSwapPointer(&((*Node)(tail).next), nil, newNode) {
			break
		}
		runtime.Gosched()
	}
	atomic.CompareAndSwapPointer(&(self.tail), tail, newNode)
}

func (self *Stack) PopStack() (val interface{}) {
	var head, tail, pretail unsafe.Pointer

	for {
		head = self.head
		tail = self.tail
		// next = ((*Node)(head)).next
		if head != self.head {
			runtime.Gosched()
			continue
		}
		pretail = self.head
		for ((*Node)(pretail)).next != self.tail {
			pretail = ((*Node)(pretail)).next
		}
		if ((*Node)(pretail)).next != self.tail {
			runtime.Gosched()
			continue
		}

		if head == tail {
			atomic.CompareAndSwapPointer(&(self.tail), tail, pretail)
		} else {
			val = ((*Node)(tail)).val
			atomic.CompareAndSwapPointer(&((*Node)(pretail)).next, tail, nil)
			atomic.CompareAndSwapPointer(&(self.tail), tail, pretail)
			return val
		}

		runtime.Gosched()
	}
}

func (self *Stack) Print() {
	p := self.head
	fmt.Printf("Stack is : ")
	for ; p != nil; p = (*Node)(p).next {
		if (*Node)(p).val != nil {
			fmt.Printf("%v ", (*Node)(p).val)
		}
	}
	fmt.Printf("\n")
}

func main() {
	q := StackNew()
	q.EnStack(5)
	q.Print()
	q.EnStack(8)
	q.Print()
	fmt.Println("POP ", q.PopStack())
	fmt.Println("POP ", q.PopStack())
	q.Print()
}
