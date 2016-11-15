package tree

import "testing"

func TestNewBinaryTree(t *testing.T) {
	tree := NewTree()
	if tree.root != nil {
		t.Fatal("root should be empty")
	}
}

func TestBinaryTreeRoot(t *testing.T) {
	tree := NewTree()
	root, err := tree.AddRoot("root")
	if err != nil {
		t.Fatal(err)
	}

	if root.Value() == "" {
		t.Fatal("root should not be empty")
	}

	if root.Left() != nil {
		t.Fatal("left children should be nil")
	}
	if root.Right() != nil {
		t.Fatal("right children should be nil")
	}
}
func TestBinaryTreeRootSet(t *testing.T) {
	tree := NewTree()
	_, err := tree.AddRoot("root")
	if err != nil {
		t.Fatal(err)
	}

	setRoot, err := tree.SetRoot("set-root")
	if err != nil {
		t.Fatal(err)
	}
	if tree.Root() != setRoot {
		t.Fatal("tree root should equal set-root")
	}
}

func TestBinaryAddLeftRight(t *testing.T) {
	tree := NewTree()
	root, err := tree.AddRoot("root")
	if err != nil {
		t.Fatal(err)
	}

	if root.Value() == "" {
		t.Fatal("root should not be empty")
	}

	left := root.AddLeft("root-left")
	right := root.AddRight("root-right")

	if root.Left() == nil {
		t.Fatal("left child of root should not be nil")
	}
	if root.Left() != left {
		t.Fatal("left child of root should equal to left node")
	}
	if root.Right() == nil {
		t.Fatal("right child of root should not be nil")
	}
	if root.Right() != right {
		t.Fatal("right child of root should equal to left node")
	}
}

func TestBinaryAddTwoTimesLeft(t *testing.T) {
	tree := NewTree()
	root, err := tree.AddRoot("root")
	if err != nil {
		t.Fatal(err)
	}

	if root.Value() == "" {
		t.Fatal("root should not be empty")
	}

	left := root.AddLeft("root-left")
	left2 := left.AddLeft("root-left2")

	if root.Left() == nil {
		t.Fatal("left child of root should not be nil")
	}
	if root.Left() != left {
		t.Fatal("left child of root should equal to left node")
	}
	if root.Left().Left() == nil {
		t.Fatal("left child of root should not be nil")
	}
	if root.Left().Left() != left2 {
		t.Fatal("left child of root should equal to left2 node")
	}
}

func TestBinaryAddTwoTimesRight(t *testing.T) {
	tree := NewTree()
	root, err := tree.AddRoot("root")
	if err != nil {
		t.Fatal(err)
	}

	if root.Value() == "" {
		t.Fatal("root should not be empty")
	}

	right := root.AddRight("root-right")
	right2 := right.AddRight("root-right2")

	if root.Right() == nil {
		t.Fatal("right child of root should not be nil")
	}
	if root.Right() != right {
		t.Fatal("right child of root should equal to right node")
	}
	if root.Right().Right() == nil {
		t.Fatal("right child of root should not be nil")
	}
	if root.Right().Right() != right2 {
		t.Fatal("right child of root should equal to right2 node")
	}
}
