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
	fmt.Printf("%+v", e)
	checkError(e)
}

type Error struct {
}

func (Error) Error() string {
	return "ooo"
}

var e Error

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
