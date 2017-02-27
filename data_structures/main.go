package main

import (
	"fmt"

	"github.com/mohamedelfiky/problem_solving_refreshment/data_structures/circular_queue"
)

func main() {
	l := circular_queue.NewCircularQueue()

	l.Push(1)
	l.Push(5)
	l.Push(6)
	l.Push(7)

	fmt.Println(l)

	fmt.Println(l.Pop(), l)

	l.Push(10)

	fmt.Println(l.Pop(), l)
	fmt.Println(l.Pop(), l)
	fmt.Println(l.Pop(), l)
	fmt.Println(l.Pop(), l)

}
