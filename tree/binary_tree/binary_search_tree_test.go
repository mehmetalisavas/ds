package tree

import "testing"

func TestSearchNil(t *testing.T) {
	tree := NewTree()
	node := Search(tree.Root(), 4)
	if node != nil {
		t.Fatal("search result should be nil")
	}
}

func TestSearchRootFetch(t *testing.T) {
	tree := NewTree()
	value := 5
	rootNode, err := tree.AddRoot(value)
	if err != nil {
		t.Fatal(err)
	}
	node := Search(rootNode, 4)
	if node != nil {
		t.Fatal("search result should be nil")
	}
	node = Search(rootNode, 5)
	if node == nil {
		t.Fatal("search result should not be nil")
	}
	if node != rootNode {
		t.Fatal("searched node should equal root node")
	}
}

//
// INSERTION TESTS
//

func TestInsertRoot(t *testing.T) {
	root := Insert(nil, 3)
	if root == nil {
		t.Fatal("root should not be nil")
	}
	if root.Value() != 3 {
		t.Fatal("root value should equal 3")
	}

	node := Insert(root, 8)
	if node == nil {
		t.Fatal("root should not be nil")
	}

	if node.Value() != 8 {
		t.Fatal("node value should equal 8")
	}
	if root.Right() != node {
		t.Fatal("right of root should equal given node")
	}
	if root.Left() != nil {
		t.Fatal("left of root should be nil")
	}
	if node.Parent() != root {
		t.Fatal("parent of node should equal root")
	}
	if node.Root() != root {
		t.Fatal("root of node should equal root")
	}

	nodeR := Insert(node, 9)
	if nodeR.Value() != 9 {
		t.Fatal("nodeR value should equal 9")
	}
	if node.Right() != nodeR {
		t.Fatal("right of node should equal ", nodeR)
	}
	if node.Left() != nil {
		t.Fatal("left of root should be nil")
	}
	if nodeR.Parent() != node {
		t.Fatal("parent of nodeR should equal node")
	}
	if root.Right().Right() != nodeR {
		t.Fatal("2*Right() of root should equal nodeR")
	}
}
