package binary_tree

type BinaryTree interface {

	CreateTree(size int,begin int) error
	DestroyTree()
	SearchNode(pos int) (*int,error)
	AddNode(where ,tree_pos,val int) error
	DeleteNode(where int)error
	PrintTree()
}

