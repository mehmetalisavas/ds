package tree

import "errors"

type Tree struct {
	root *Node
}

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	value  interface{}
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
	// parent of root should be parent
	t.root.setParent(t.root)

	return t.root, nil
}

// SetRoot sets the root, if any value is set for root, it overrides the previous value
func (t *Tree) SetRoot(val interface{}) *Node {
	t.root.value = val

	return t.root
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

// Right exports the unexported right node
func (n *Node) Parent() *Node { return n.parent }

// Root fetches the root node of the tree
func (n *Node) Root() *Node {
	for n.parent != n {
		n = n.Parent()
	}

	return n
}

func (n *Node) AddLeft(val interface{}) *Node {
	return n.addLeft(val)
}

func (n *Node) AddRight(val interface{}) *Node {
	return n.addRight(val)
}

// IsLeaf checks the node if its leaf of tree
func (n *Node) IsLeaf() bool {
	return n.value != nil && n.Left() == nil && n.Right() == nil
}

func (n *Node) addLeft(val interface{}) *Node {
	n.left = newNode(val)
	n.left.setParent(n)

	return n.left
}

func (n *Node) addRight(val interface{}) *Node {
	n.right = newNode(val)
	n.right.setParent(n)

	return n.right
}
func (n *Node) setParent(node *Node) {
	n.parent = node
}
