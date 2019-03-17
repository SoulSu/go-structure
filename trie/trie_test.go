package trie

import (
	"fmt"
	"testing"
)

func Init() {

}

func TestTrie_Add(t *testing.T) {
	te := NewTrie()
	te.Add("hello")

	fmt.Println(te.Container("hell"))
	fmt.Println(te.Container("hello"))
	fmt.Println(te.Container("hello"))
}
