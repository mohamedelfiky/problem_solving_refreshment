package main

import (
	"fmt"

	"github.com/mohamedelfiky/problem_solving_refreshment/data_structures/priority_queue"
)

func main() {
	l := new(priority_queue.PriorityQueue)

	l.Push(1, func() { fmt.Println("Execute 1") })
	l.Push(5, func() { fmt.Println("Execute 5") })
	l.Push(15, func() { fmt.Println("Execute 15") })
	l.Push(6, func() { fmt.Println("Execute 6") })
	l.Push(7, func() { fmt.Println("Execute 7") })
	l.Push(70, func() { fmt.Println("Execute 70") })
	l.Push(3, func() { fmt.Println("Execute 3") })

	fmt.Println(l)
	fmt.Println(l.Size())
	l.Pop().Operation()
	l.Pop().Operation()
	l.Pop().Operation()
	l.Pop().Operation()
	l.Pop().Operation()
	l.Pop().Operation()
	l.Pop().Operation()

	fmt.Println(l)
	fmt.Println(l.Size())
}
