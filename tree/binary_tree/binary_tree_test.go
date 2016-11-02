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
    root,err := tree.AddRoot("root")
    if err != nil {
        t.Fatal(err)
    }

    if root.Value() == ""{
        t.Fatal("root should not be empty")
    }

    if root.Left() != nil {
        t.Fatal("left children should be nil")
    }
    if root.Right() != nil {
        t.Fatal("right children should be nil")
    }
}

func TestBinaryAddLeftRight(t *testing.T) {
    tree := NewTree()
    root,err := tree.AddRoot("root")
    if err != nil {
        t.Fatal(err)
    }

    if root.Value() == ""{
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
