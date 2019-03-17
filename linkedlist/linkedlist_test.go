package linkedlist

import (
	"testing"
)

func Init() *LinkedList {
	l := NewLinkedList()
	l.AddElement(1)
	l.AddElement(9)
	l.AddElement(8)
	l.AddElement(7)
	l.AddElement(666)
	return l
}

func TestLinkedList_String(t *testing.T) {
	l := Init()
	l.String()
}

func TestLinkedList_RemoveElement(t *testing.T) {
	l := Init()
	l.String()
	l.RemoveElement(8)
	l.String()
}
