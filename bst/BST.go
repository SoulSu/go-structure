// 二叉搜索树
// 1. 左节点的值小于右节点
// 2. 最坏情况下退化成一个链表
// 3. 最小值是在最左边的节点
// 4. 最大值是在最右边的节点

package bst

import "fmt"

type BST struct {
	root *Node // 根节点
	size int   // 大小
}

type Node struct {
	E     int // 使用int作为元素值
	left  *Node
	right *Node
}

func NewBST() *BST {
	return NewBSTWithNode(nil)
}

func NewBSTWithNode(node *Node) *BST {
	bst := &BST{root: node}
	if node != nil {
		bst.size++
	}
	return bst
}

// 向二叉搜索树中插入一个节点
func (bst *BST) AddElement(e int) {
	node := &Node{e, nil, nil}
	// 还没有任何元素
	if bst.root == nil {
		bst.root = node
		return
	}

	bst.addElement(bst.root, node)

}

func (bst *BST) addElement(root, node *Node) {
	if node.E < root.E && root.left == nil {
		root.left = node
		return
	} else if node.E > root.E && root.right == nil {
		root.right = node
		return
	}

	if node.E < root.E {
		bst.addElement(root.left, node)
	} else if node.E > root.E {
		bst.addElement(root.right, node)
	}
}

func (bst *BST) Container(e int) *Node {
	return bst.container(bst.root, e)
}

func (bst *BST) container(root *Node, e int) *Node {
	if root == nil {
		return nil
	}

	if e == root.E {
		return root
	}

	if e > root.E {
		bst.container(root.right, e)
	}
	if e < root.E {
		bst.container(root.left, e)
	}

	return nil
}

// 删除最小节点
func (bst *BST) DeleteMinElement() {
	bst.root = bst.deleteMinElement(bst.root)
}

// 删除最小节点 返回根节点
func (bst *BST) deleteMinElement(root *Node) *Node {

	if root == nil {
		return nil
	}

	// 这个就是要删除的节点
	if root.left == nil {
		root = nil
		return nil
	}

	root.left = bst.deleteMinElement(root.left)

	return root
}

func (bst *BST) DeleteMaxElement() {
	bst.root = bst.deleteMaxElement(bst.root)
}

func (bst *BST) deleteMaxElement(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.right == nil {
		root = nil
		return nil
	}

	root.right = bst.deleteMaxElement(root.right)

	return root
}

func (bst *BST) findMin(root *Node) *Node {
	for root.left != nil {
		root = root.left
	}

	return root
}

func (bst *BST) findMax(root *Node) *Node {
	for root.right != nil {
		root = root.right
	}

	return root
}

func (bst *BST) removeElement(root *Node, e int) *Node {

	if root == nil {
		return nil
	}

	if root.E > e {
		root.left = bst.removeElement(root.left, e)
		return root
	} else if root.E < e {
		root.right = bst.removeElement(root.right, e)
		return root
	} else {
		if root.left == nil {
			removeEle := root.right
			root.right = nil
			return removeEle
		}
		if root.right == nil {
			removeEle := root.left
			root.left = nil
			return removeEle
		}

		// 左右节点都不为空的情况
		// 1. 找到删除节点后面的最小节点
		// 2. 删除最小节点
		// 3. 最小节点代替删除节点的位置
		minNode := bst.findMin(root)
		minNode.left = bst.deleteMinElement(root)
		minNode.right = root.right
		return minNode
	}

	panic("不会运行到这里来的")
}

//              9
//       5          15
//    1    6      10   16
//  0        7
// 1. 如果删除的是最下面的叶子节点,那么直接将该节点父节点的指向为空nil即可
// 2. 如果删除的是有只有左右节点的根节点,那么只需要将左右节点的前驱或者后继节点接上父节点的删除位置指向
// 3.
func (bst *BST) RemoveElement(e int) *Node {
	return bst.removeElement(bst.root, e)
}

func (bst *BST) Size() int {
	return bst.size
}

func (bst *BST) IsEmpty() bool {
	return bst.size == 0
}

func (bst *BST) String() {

}

// 中序遍历
func (bst *BST) MidIter() {
	bst.midIter(bst.root)
}

func (bst *BST) midIter(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.E)
	bst.midIter(root.left)
	bst.midIter(root.right)
}
