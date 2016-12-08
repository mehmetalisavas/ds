package tree

// Search searches the given integer value inside the binary tree
// If given value is greater than root's value, then search right nodes of tree
// otherwise search in the left nodes of the root node
func Search(root *Node, value int) *Node {
	if root == nil || root.Value().(int) == value {
		return root
	}

	if value > root.Value().(int) {
		Search(root.Right(), value)
	}

	return Search(root.Left(), value)
}
