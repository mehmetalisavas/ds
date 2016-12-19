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

// Insert insert the value into the BST and returns last added node
func Insert(node *Node, value int) *Node {
	if node == nil {
		node = newNode(value)
		node.setParent(node)
		return node
	}

	if value > node.Value().(int) {
		node.right = Insert(node.Right(), value)
		node.right.setParent(node)
		return node.right
	}

	node.left = Insert(node.Left(), value)
	node.left.setParent(node)
	return node.left
}
