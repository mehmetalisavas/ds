package tree

import "errors"

type Tree struct {
	root *Node
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
	}
}

// Root specifies the root of tree
func (t *Tree) Root() *Node { return t.root }

// AddRoot sets the the if any value is not set
func (t *Tree) AddRoot(val interface{}) (*Node, error) {
	if t.root != nil {
		return nil, ErrAlreadyHasRoot
	}

	t.root = newNode(val)

	return t.root, nil
}

// SetRoot sets the root, if any value is set for root, it overrides the previous value
func (t *Tree) SetRoot(val interface{}) (*Node, error) {
	t.root.value = val

	return t.root, nil
}

func newNode(val interface{}) *Node {
	return &Node{
		value: val,
		left:  nil,
		right: nil,
	}
}

// Value exports the unexported node value
func (n *Node) Value() interface{} { return n.value }

// Left exports the unexported left node
func (n *Node) Left() *Node { return n.left }

// Right exports the unexported right node
func (n *Node) Right() *Node { return n.right }

func (n *Node) AddLeft(val interface{}) *Node {
	return n.addLeft(val)
}

func (n *Node) AddRight(val interface{}) *Node {
	return n.addRight(val)
}

func (n *Node) addLeft(val interface{}) *Node {
	n.left = newNode(val)

	return n.left
}

func (n *Node) addRight(val interface{}) *Node {
	n.right = newNode(val)

	return n.right
}
