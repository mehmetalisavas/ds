package tree

// Search searches the given integer value inside the binary tree
// If given value is greater than root's value, then search right nodes of tree
// otherwise search in the left nodes of the root node
func Search(root *Node, value int) *Node {
	if root == nil || root.Value().(int) == value {
		return root
	}

	if value > root.Value().(int) {
		return Search(root.Right(), value)
	}

	return Search(root.Left(), value)
}

// Insert insert the value into the BST
// TODO ~mehmetali
// add tests
func Insert(node *Node, value int) *Node {
	if node == nil {
		return newNode(value)
	}

	if value > node.Value().(int) {
		Insert(node.Right(), value)
	} else {
		Insert(node.Left(), value)
	}

	return node
}
