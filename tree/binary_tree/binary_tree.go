package tree

import "errors"

type Tree struct {
	root *Node
	size int
}

type Node struct {
	left  *Node
	right *Node
	value interface{}
}

var (
	ErrAlreadyHasRoot = errors.New("Tree already has a root")
)

func NewTree() *Tree {
	return &Tree{
		root: nil,
		size: 0,
	}
}

func (t *Tree) Root(val interface{}) (*Tree, error) {
	if t.root != nil {
		return nil, ErrAlreadyHasRoot
	}

	t.root = NewNode(val)
	t.size++

	return t, nil
}

func NewNode(val interface{}) *Node {
	return &Node{
		value: val,
		left:  nil,
		right: nil,
	}
}

func (n *Node) Left(val interface{}) *Node {
	n.left = NewNode(val)

	return n.left
}

func (n *Node) Right(val interface{}) *Node {
	n.right = NewNode(val)

	return n.right
}
