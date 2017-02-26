package main

import (
	"fmt"

	"github.com/mohamedelfiky/problem_solving_refreshment/data_structures/doubly_linked_list"
)

func main() {
	l := new(doubly_linked_list.DoublyLinkedList)

	l.Push(1)
	l.Push(5)
	n2 := l.Push(10)
	n3 := l.Push(15)
	l.Remove(n2)

	l.Remove(n3)
	fmt.Println(l, n2, n3)
}
