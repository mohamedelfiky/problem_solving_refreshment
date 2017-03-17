package main

import (
	"fmt"

	"github.com/mohamedelfiky/problem_solving_refreshment/data_structures/hashmap"
)

func main() {
	l := hashmap.New()

	l.Put("key1", "45678")
	l.Put("key2", "4577")
	fmt.Println(l)
	l.Remove("key2")
	l.Remove("key1")
	fmt.Println(l.Get("key1"))
	fmt.Println(l.Get("key2"))
}
