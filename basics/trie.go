package main

import (
	"fmt"

	"github.com/dghubble/trie"
)

func main() {
	fmt.Println("Hello, playground")
	t := trie.NewPathTrie()
	fmt.Println(t)
	t.Put("/foo", 1)
}
