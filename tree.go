package main

import (
    "fmt"
    "os"
    "math/rand"
    "github.com/tools"
    "time"
)

//--------------------------------tree basic opereation--------------------------------------//
type Tree struct{
    Left  *Tree
    Value int
    Right *Tree
}

func Insert(t *Tree, v int) *Tree{
    if t == nil {
        return &Tree{nil, v, nil}
    }
    if v < t.Value {
        t.Left = Insert(t.Left, v)
    } else {
        t.Right = Insert(t.Right, v)
    }
    return t
}

func New(k int) *Tree{
    var t *Tree
    for _,v := range rand.Perm(10) {
        fmt.Println(v)
        t = Insert(t, (1+v)*k)
    }
    return t
}


func (t *Tree) String() string{
    if t == nil {
        return "()"
    }
    s := ""
    if t.Left != nil {
        // Pay attention: fmt sentence's parameter can not recursive.
        // s = fmt.Sprintf("%s ", t.Left.String())
        s += t.Left.String() + " "
    }
    s += fmt.Sprintf("%d", t.Value)
    if t.Right != nil {
        // s = fmt.Sprintf(" %s", t.Right.String())
        s += " " + t.Right.String()
    }
    return "(" + s + ")"
}

func (t *Tree) Len() int {
    var lh, rh, h int
    if t == nil {
        return 0
    }
    if t.Left != nil {
        lh = t.Left.Len()
    } 
    if t.Right != nil {
        rh = t.Right.Len()
    }
    if lh > rh {
        h = lh + 1
    } else {
        h = rh + 1
    }
    return h
}

func (t *Tree) Distance() int {
    // calculate for the farest path
    fmt.Println("Distance......")
    stack := tools.StackNew()
    stack2 := tools.StackNew()
    if t == nil {
        return 0
    }

    stack.EnStack(t)
    for {
        q := stack.PopStack()
        if q == nil {
            break
        }
        p := q.(*Tree)
        stack2.EnStack(p)
        fmt.Println("in stack by tree ", p.Value)
        if p.Left != nil {
            stack.EnStack(p.Left)       
        }
        if p.Right != nil {
            stack.EnStack(p.Right)
        }
    }
    //calculate the max deep one by one
    max := 0
    for {
        t := stack2.PopStack()
        if t == nil {
            return max
        }
        s := t.(*Tree)
        d := 0
        if s.Left != nil {
            d = d + s.Left.Len() + 1
        }
        if s.Right != nil {
            d = d + s.Right.Len() + 1
        }
        d = d-1
        if d > max {
            max = d
        }

    }
    
    return max
}

func (t *Tree) Walk(c chan int) {
    if t == nil {
        close(c)
        return
    }
    c <- t.Value
    if t.Left != nil {
        go t.Left.Walk(c)
    }
    if t.Right != nil {
        go t.Right.Walk(c)
    }
}


//-------------------------------------------------------------------------------------------//
func main() {
    t := New(1)
    fmt.Fprintf(os.Stderr, "Tree is %s\n", t.String())

    // fmt.Println(t.Len())
    // fmt.Println("The max path is ", t.Distance())
    c := make(chan int)
    go t.Walk(c)
    var v int
    // go func(){
            for {
            select {
            case v=<-c:
                fmt.Print(v, " | ")

            }
        }
    // }()
    time.Sleep(time.Second*10)
}
