package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	rangeTree(t, ch)
	close(ch)
}

func rangeTree(t *tree.Tree, ch chan int) {
	if t != nil {
		//fmt.Printf("tree %v left is %v\n", t, t.Left)
		rangeTree(t.Left, ch)
		//fmt.Printf("tree value is %d ----------\n", t.Value)
		ch <- t.Value
		//fmt.Printf("tree %v right is %v\n", t, t.Right)
		rangeTree(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := range c1 {
		if i != <-c2 {
			return false
		}
	}
	return true
}

func main() {
	// create random bTree
	t1 := tree.New(1)
	t2 := tree.New(1)

	if Same(t1, t2) == true {
		fmt.Printf("t1 %s is equal to t2 %s\n", t1.String(), t2.String())
	} else {
		fmt.Printf("t1 %s is not equal to t2 %s\n", t1.String(), t2.String())
	}

}
