// 链表
package linkedlist

import "fmt"

type LinkedList struct {
	root *Node
	size int
}

type Node struct {
	E    int
	next *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func NewLinkedListWithElement(e int) *LinkedList {
	return &LinkedList{
		size: 1,
		root: &Node{e, nil},
	}
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// 直接将新元素放在链表最前面
func (l *LinkedList) AddElement(e int) {
	l.root = l.addFirst(l.root, e)
	l.size++
}

func (l *LinkedList) addFirst(root *Node, e int) *Node {
	root = &Node{e, root}
	return root
}

func (l *LinkedList) addLast(root *Node, e int) *Node {

	if root.next == nil {
		root.next = &Node{e, nil}
		return root
	}

	//
	root.next = l.addLast(root.next, e)
	return root
}

// 查找链表如果找到返回node != nil
func (l *LinkedList) Container(e int) *Node {
	return l.container(l.root, e)
}

func (l *LinkedList) container(root *Node, e int) *Node {
	if root == nil {
		return nil
	}
	// 找到了
	if root.E == e {
		return root
	}
	// 继续往下面找
	return l.container(root.next, e)
}

func (l *LinkedList) RemoveElement(e int) {
	root := l.removeElement(l.root, e)
	if root != nil {
		l.root = root
		l.size--
	}
}

func (l *LinkedList) removeElement(root *Node, e int) *Node {
	if root == nil {
		return nil
	}

	// 找到要删除的节点了
	if root.E == e {
		next := root.next
		root = nil
		return next
	}

	root.next = l.removeElement(root.next, e)

	return root
}

// 打印链表
func (l *LinkedList) String() {
	node := l.root
	for node != nil {
		fmt.Printf("%d -> ", node.E)
		node = node.next
	}
	fmt.Println()
}
