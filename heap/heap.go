package heap

import "fmt"

// 最大堆
type MaxHeap struct {
	items []int
	size  int
	cap   int
}

func NewMaxHeap(cap int) *MaxHeap {
	return &MaxHeap{
		items: make([]int, cap+1),
		size:  0,
		cap:   cap,
	}
}

func (m *MaxHeap) Size() int {
	return m.size
}

func (m *MaxHeap) IsEmpty() int {
	return m.size == 0
}

// 获取父节点的坐标
func (m *MaxHeap) getParent(i int) int {
	return i / 2
}

// 获取左子树的坐标
func (m *MaxHeap) getLeft(i int) int {
	return i * 2
}

// 获取右子树的坐标
func (m *MaxHeap) getRight(i int) int {
	return i*2 + 1
}

// 向上调整堆性质
func (m *MaxHeap) shiftUp(i int) {
	// i节点比自己父节点值大就交换
	for i > 1 && m.items[i] > m.items[m.getParent(i)] {
		m.items[i], m.items[m.getParent(i)] = m.items[m.getParent(i)], m.items[i]
		i = m.getParent(i)
	}
}

func (m *MaxHeap) Add(t int) {

	if m.size+1 >= m.cap {
		return
	}

	// 加到最后面
	m.items[m.size+1] = t
	m.size++

	// 检查t是否满足堆性质
	m.shiftUp(m.size)
}

func (m *MaxHeap) Print() {
	fmt.Println(m.items)
}

// 向下检查是否满足最大堆性质
func (m *MaxHeap) shiftDown(i int) {
	// i节点比自己节点值大就交换
	for 2*i < m.size {

		l := m.getLeft(i)
		r := m.getRight(i)
		j := l
		if r <= m.size && m.items[l] < m.items[r] {
			j = r
		}

		if m.items[i] > m.items[j] {
			break
		}
		m.items[i], m.items[j] = m.items[j], m.items[i]

		i = j
	}
}

func (m *MaxHeap) PopMax() int {

	if m.size-1 <= 0 {
		return 0
	}

	max := m.items[1]

	m.items[1] = m.items[m.size]
	m.items[m.size] = -1
	m.size--

	m.shiftDown(1)

	return max
}
