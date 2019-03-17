// 二叉搜索树
// 1. 左节点的值小于右节点
// 2. 最坏情况下退化成一个链表
// 3. 最小值是在最左边的节点
// 4. 最大值是在最右边的节点

package bst

import (
	"fmt"
	"testing"
)

func Init() *BST {
	bst := NewBST()

	bst.AddElement(10)
	bst.AddElement(9)
	bst.AddElement(6)
	bst.AddElement(13)
	bst.AddElement(17)

	return bst
}

func TestBST_midIter(t *testing.T) {
	bst := Init()
	bst.MidIter()
}

func TestBst_findMin(t *testing.T) {
	bst := Init()

	fmt.Println(bst.findMin(bst.root))
}

func TestBst_findMax(t *testing.T) {
	bst := Init()

	fmt.Println(bst.findMax(bst.root))
}

func TestBST_RemoveElement(t *testing.T) {
	bst := Init()
	bst.RemoveElement(6)
	fmt.Println("-------------------")
	bst.MidIter()
}
