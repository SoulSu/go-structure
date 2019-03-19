package heap

import (
	"testing"
)

func TestMaxHeap_Add(t *testing.T) {
	h := NewMaxHeap(10)

	h.Add(5)
	h.Print()
	h.Add(8)
	h.Add(9)
	h.Print()
	h.Add(7)
	h.Add(1)
	h.Add(9)
	h.Add(10)
	h.Print()
	h.Add(19)
	h.Print()

	h.PopMax()
	h.Print()
	h.PopMax()
	h.Print()
}
