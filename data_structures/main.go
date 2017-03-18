package main

import (
	"fmt"

	"github.com/mohamedelfiky/problem_solving_refreshment/data_structures/trie"
)

func main() {
	l := trie.NewTrie()

	l.Insert("key")
	l.Insert("jojo")
	l.Insert("koz")
	fmt.Println(l.WordExist("key"))
	fmt.Println(l.WordExist("jojo"))
	fmt.Println(l.PrefixExist("joj"))
	fmt.Println(l.PrefixExist("jof"))
}
