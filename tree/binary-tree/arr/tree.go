package main

import (
	"errors"
	"math"
	"fmt"
	"go-structure/tree/binary-tree"
)



// 使用数组表示二叉树
//               A
//            B    C
//           D E  F G
//
//       []int 0 1 2 3 4 5 6
//       []int A B C D E F G
//       pos_left = index*2 + 1
//       pos_right = index*2 + 2
//
type MyTree struct {
	tree []*int
	size int
}

func (this *MyTree) CreateTree(size int, begin int) error {
	this.tree = make([]*int, size, size)
	this.tree[0] = &begin
	this.size = size
	return nil
}

func (this *MyTree) DestroyTree() {
	this.tree = nil
}

func (this *MyTree) SearchNode(pos int) (*int, error) {
	node := this.tree[pos]
	if this.size < pos {
		return nil, errors.New("pos err")
	}
	if *node == 0 {
		return nil, errors.New("pos val empty")
	}
	return node, nil
}

func (this *MyTree)AddNode(where, tree_pos, val int) error {

	if where > this.size {
		return errors.New("size to long")
	}

	if this.tree[where] == nil {
		return errors.New("pos is not empty")
	}

	// add left
	if tree_pos == 0 {
		if where * 2 + 1 > this.size {
			return errors.New("can not add left")
		}
		if this.tree[where * 2 + 1] != nil && *this.tree[where * 2 + 1] != 0 {
			return errors.New("ps left is not empty")
		}
		this.tree[where * 2 + 1] = &val
	}
	// add right
	if tree_pos == 1 {
		if where * 2 + 2 > this.size {
			return errors.New("can not add right")
		}
		if this.tree[where * 2 + 2] != nil &&*this.tree[where * 2 + 2] != 0 {
			return errors.New("ps right is not empty")
		}
		this.tree[where * 2 + 2] = &val
	}

	return nil
}

func (this *MyTree)DeleteNode(where int) error {
	if this.size < where {
		return errors.New("tooooooooo longggggg")
	}

	if *this.tree[where] == 0 {
		return errors.New("not delete")
	}
	zero := 0
	this.tree[where] = &zero

	return nil
}
func (this *MyTree) p(i int) int {
	if this.tree[i] == nil {
		return 0
	} else {
		return *this.tree[i]
	}
}
func (this *MyTree) PrintTree() {
	//log.Println(this.tree)
	//1, 2, 4, 8, 16   2 ^ i
	deep := 0
	p := 0
	for i := 0; i < this.size; i++ {
		p++
		fmt.Print(this.p(i))
		if float64(p) == math.Pow(2, float64(deep)) {
			p = 0
			deep++
			fmt.Println()
		}
	}

}

func main() {

	tree := new(MyTree)

	i := (binary_tree.BinaryTree)(tree)

	fmt.Println(tree == i)
	i.DestroyTree()

	tree.CreateTree(20, 1)

	tree.AddNode(0, 0, 2)
	tree.AddNode(0, 1, 3)
	tree.AddNode(1, 1, 4)
	tree.AddNode(1, 0, 5)
	tree.AddNode(2, 0, 6)
	tree.AddNode(2, 1, 7)

	tree.AddNode(6, 0, 13)

	tree.PrintTree()
}







