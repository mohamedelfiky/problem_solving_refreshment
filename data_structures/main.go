package main

import (
	"fmt"

	"github.com/mohamedelfiky/problem_solving_refreshment/data_structures/heap"
)

func main() {
	l := new(heap.MinHeap)

	l.Push(1)
	l.Push(5)
	l.Push(15)
	l.Push(6)
	l.Push(7)
	l.Push(70)
	l.Push(3)

	fmt.Println(l)
	fmt.Println(l.Size())
	l.Pop()

	fmt.Println(l)
	fmt.Println(l.Size())
}
