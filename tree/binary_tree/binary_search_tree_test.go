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
