package main

import "fmt"

// interview question
// TreeNode represents a single node in the binary tree
type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

// Insert adds a new node to the binary tree
func (n *TreeNode) Insert(v int) {
    if v <= n.Value {
        if n.Left == nil {
            n.Left = &TreeNode{Value: v}
        } else {
            n.Left.Insert(v)
        }
    } else {
        if n.Right == nil {
            n.Right = &TreeNode{Value: v}
        } else {
            n.Right.Insert(v)
        }
    }
}

// Search looks for a value in the binary tree
func (n *TreeNode) Search(v int) bool {
    if n == nil {
        return false
    }

    if n.Value == v {
        return true
    } else if v < n.Value {
        return n.Left.Search(v)
    } else {
        return n.Right.Search(v)
    }
}

// InOrderTraversal prints the tree in sorted ascending order
func (n *TreeNode) InOrderTraversal() {
    if n == nil {
        return
    }
    n.Left.InOrderTraversal()
    fmt.Printf("%d ", n.Value)
    n.Right.InOrderTraversal()
}

func main() {
    // Create the root node
    root := &TreeNode{Value: 10}

    // Insert values
    root.Insert(5)
    root.Insert(15)
    root.Insert(3)
    root.Insert(7)
    root.Insert(12)
    root.Insert(18)

    // Print the tree in sorted order
    fmt.Print("InOrder Traversal: ")
    root.InOrderTraversal()
    fmt.Println()

    // Search for a value
    fmt.Println("Search 7:", root.Search(7))
    fmt.Println("Search 13:", root.Search(13))
}
