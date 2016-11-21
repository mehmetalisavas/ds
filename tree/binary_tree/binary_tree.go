package tree

import "errors"

// Tree specifies only root.
// then ready to go for node after defining tree
type Tree struct {
	root *Node
}

// Node specifies the preferences of each node
type Node struct {
	parent *Node
	left   *Node
	right  *Node
	value  interface{}
}

// binary tree errors
var (
	ErrAlreadyHasRoot = errors.New("Tree already has a root")
)

// NewTree inits the Tree
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

// Parent exports the unexported parent value
func (n *Node) Parent() *Node { return n.parent }

// Root fetches the root node of the tree
func (n *Node) Root() *Node {
	for n.parent != n {
		n = n.Parent()
	}

	return n
}

// AddLeft add a child to the left side of node
func (n *Node) AddLeft(val interface{}) *Node {
	return n.addLeft(val)
}

// AddRight add a child to the right side of node
func (n *Node) AddRight(val interface{}) *Node {
	return n.addRight(val)
}

// IsLeaf checks the node if its leaf of tree
func (n *Node) IsLeaf() bool {
	return n.value != nil && n.Left() == nil && n.Right() == nil
}

// IsRoot checks the node if its root
func (n *Node) IsRoot() bool {
	return n == n.Parent()
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
