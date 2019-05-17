package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := []int{5, 8, 9, 4, 3, 2, 0, 1, 7, 6}

	fmt.Printf("Original a is %v\n", a)
	// bubblesort(a)
	insertsort(a)
	// quicksort(a)
	// heapsort(a, 0, len(a))
	fmt.Printf("Sorted a is %v\n", a)
}

//-------------------------------------------------------
func heapsort(data []int, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	fmt.Printf("first is %d, lo is %d, hi is %d\n", first, lo, hi)
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
		fmt.Printf("After sift down(%d,%d,%d) %v\n", i, hi, first, data)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		data[first], data[first+i] = data[first+i], data[first]
		fmt.Printf("After swap %d and %d, %v\n", data[first+i], data[first], data)
		siftDown(data, lo, i, first)
		fmt.Printf("After sift down(%d,%d,%d) %v\n", lo, i, first, data)
	}
	fmt.Println("---------------------------------------------------")
}

// siftDown implements the heap property on data[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDown(data []int, lo, hi, first int) {
	root := lo
	for {
		// child is the index of root's left child.
		child := 2*root + 1
		// if the child index is out of tree range, break.
		if child >= hi {
			break
		}
		// if the right child index is still in the tree range, compare the left child and the right child
		if child+1 < hi && data[first+child] < data[first+child+1] {
			child++// if the left child is less than the right child, point to the right child else still point the left child
		}
		// if the root is not less than the bigger child, return
		if !(data[first+root] < data[first+child]) {
			return
		}
		// Swap the bigger child and the root, which may be right child or left child.
		data[first+root], data[first+child] = data[first+child], data[first+root]
		root = child
	}
}

//---------------------------------------------------------

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)
	fmt.Printf("left %d, right %d, pivot %d\n", left, right, pivot)

	a[pivot], a[right] = a[right], a[pivot]
	fmt.Printf("swap pivot and right: %v\n", a)
	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	fmt.Printf("swap each a element that is smaller than right: %v\n", a)

	a[left], a[right] = a[right], a[left]
	fmt.Printf("swap left and right: %v\n", a)
	fmt.Println("---------------------------")
	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func insertsort(a []int) {
	fmt.Println("Start with ", a)

	//---------------------------------------------------------------------
	// for i := 0; i < len(a); i++ {
	//   j := 0
	//   temp := a[i]
	//   for ; j < i; j++ {
	//     if a[j] > a[i] {
	//       // insert a[i] into array b at the right position
	//       fmt.Printf("Found b[%d] %d bigger than a[%d] %d\n", j, a[j], i,  a[i])
	//       for m := i; m>j; m--{
	//         fmt.Printf("set a[%d] tobe a[%d] %d\n", m, m-1,  a[m-1])
	//         a[m] = a[m-1]
	//       }
	//       break
	//     }
	//   }
	//   fmt.Printf("set b[%d] to be a[%d] %d\n", j, i, a[i])
	//   a[j] = temp
	// }
	// The following codes equal to the upper codes that are eaiser to understand.
	for i := 0; i < len(a); i++ {
		temp := a[i]
		j := i
		for ; j >= 1 && temp < a[j-1]; j-- {
			fmt.Printf("a[%d] %d is less than a[%d] %d\n", i, temp, j, a[j-1])
			a[j] = a[j-1]
			fmt.Println(a)
		}
		a[j] = temp
		fmt.Println(a)
		fmt.Println("===========================")
	}

	//-----------------------------------------------------------------------
	fmt.Println(a)

}

func bubblesort(a []int) {
	fmt.Println("Start with ", a)
	l := len(a)
	if l == 0 {
		fmt.Println("array is none")
		return
	}
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if a[i] > a[j] {
				fmt.Printf("swap a[%d] and a[%d]\n", i, j)
				a[i], a[j] = a[j], a[i]
				fmt.Println(a)
			}
		}
		fmt.Println("========================")
	}
	fmt.Println(a)
}
