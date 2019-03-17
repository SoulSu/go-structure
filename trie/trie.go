// 前缀树, 字典树
// 常用于单词查找
package trie

import (
	"fmt"
	"strings"
)

type Trie struct {
	root *Node
	size int
}

type Node struct {
	isWord bool // 到该节点是否是一个单词
	next   map[string]*Node
}

func NewTrie() *Trie {
	return &Trie{root: &Node{false, make(map[string]*Node)}, size: 0}
}

func (t *Trie) Size() int {
	return t.size
}

func (t *Trie) IsEmpty() bool {
	return t.size == 0
}

func (t *Trie) Add(s string) {
	ss := strings.Split(s, "")
	cur := t.root
	for _, c := range ss {
		if next, has := cur.next[c]; has {
			cur = next
		} else {
			node := &Node{false, make(map[string]*Node)}
			cur.next[c] = node
			cur = node
		}
	}

	if !cur.isWord {
		cur.isWord = true
		t.size++
	}
	fmt.Println(t.root, cur)
}

func (t *Trie) Container(s string) bool {
	ss := strings.Split(s, "")
	cur := t.root

	for _, c := range ss {
		if next, has := cur.next[c]; has {
			cur = next
		} else {
			return false
		}
	}

	return cur.isWord
}
