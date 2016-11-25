package tree

import (
	"fmt"
	"testing"
)

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
func TestBinaryTreeAddRoot(t *testing.T) {
	tree := NewTree()
	root, err := tree.AddRoot("root")
	if err != nil {
		t.Fatal(err)
	}

	if root.Value() == "" {
		t.Fatal("root should not be empty")
	}
	if root.Value() != "root" {
		t.Fatal("root value should equal 'root'")
	}
}

func TestBinaryTreeRootSet(t *testing.T) {
	tree := NewTree()
	_, err := tree.AddRoot("root")
	if err != nil {
		t.Fatal(err)
	}

	setRoot := tree.SetRoot("set-root")
	if tree.Root() != setRoot {
		t.Fatal("tree root should equal set-root")
	}
}

func TestBinaryTreeIsRoot(t *testing.T) {
	tree := NewTree()
	root, err := tree.AddRoot("root")
	if err != nil {
		t.Fatal(err)
	}

	if root.IsRoot() != true {
		t.Fatal("isRoot should be true for root value")
	}

	left := root.AddLeft("left")
	if left.IsRoot() != false {
		t.Fatal("isRoot should be true for left value")
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
func TestBinaryTreeParent(t *testing.T) {
	tree := NewTree()
	root, err := tree.AddRoot("root")
	if err != nil {
		t.Fatal(err)
	}

	if root.Parent() != root {
		t.Fatal("parent of root should equal ")
	}
}

func TestBinaryTreeParentNodes(t *testing.T) {
	tree := NewTree()
	root, err := tree.AddRoot("root")
	if err != nil {
		t.Fatal(err)
	}

	left := root.AddLeft("root-left")
	left2 := left.AddLeft("root-left2")

	if left.Parent() != root {
		t.Fatal("parent of left should equal root")
	}
	if left2.Parent() != left {
		t.Fatal("parent of left should equal root")
	}
}

func TestBinaryTreeRootOfNodes(t *testing.T) {
	tree := NewTree()
	root, err := tree.AddRoot("root")
	if err != nil {
		t.Fatal(err)
	}

	left := root.AddLeft("root-left")
	left2 := left.AddLeft("root-left2")
	left3 := left.AddLeft("root-left3")

	if left2.Root() != root {
		t.Fatal("root of left2 should equal root")
	}

	if left3.Root() != root {
		t.Fatal("root of left3 should equal root")
	}
}

func TestBinaryTreeLeaves(t *testing.T) {
	tree := NewTree()
	root, err := tree.AddRoot("root")
	if err != nil {
		t.Fatal(err)
	}

	left := root.AddLeft("root-left")
	left2 := left.AddLeft("root-left2")
	right := root.AddRight("root-right")

	if root.IsLeaf() == true {
		t.Fatal("isLeaf should equal false")
	}
	if left.IsLeaf() == true {
		t.Fatal("isLeaf should equal false")
	}

	if left2.IsLeaf() != true {
		t.Fatal("isLeaf should equal true")
	}

	if right.IsLeaf() != true {
		t.Fatal("isLeaf should equal true")
	}
}

func TestBinaryTreeRemoveIfLeaf(t *testing.T) {
	tree := NewTree()
	root, err := tree.AddRoot("root")
	if err != nil {
		t.Fatal(err)
	}

	left := root.AddLeft("root-left")
	left2 := left.AddLeft("root-left2")
	// right := root.AddRight("root-right")

	if root.IsLeaf() == true {
		t.Fatal("isLeaf should equal false")
	}
	if left.IsLeaf() == true {
		t.Fatal("isLeaf should equal false")
	}

	if root.Left().Left() != left2 {
		t.Fatal("should equal 'left2'")
	}
	fmt.Println("root.left.left", root.Left().Left())
	left2.RemoveIfLeaf()
	fmt.Println("root.left.left", root.Left().Left())
	if root.Left().Left() != nil {
		t.Fatal("should equal nil")
	}

}
