package iterator

import "fmt"

// Node of a binary tree
type Node struct {
	Value       int
	Left, Right *Node
}

// Iterator interface (generic over tree nodes)
type TreeIterator interface {
	HasNext() bool
	Next() *Node
}

// Concrete Aggregate
type BinaryTree struct {
	Root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{Root: root}
}

func (t *BinaryTree) Iterator() TreeIterator {
	it := &InOrderIterator{stack: []*Node{}}
	it.pushLeft(t.Root)
	return it
}

// Concrete Iterator: in‑order traversal
type InOrderIterator struct {
	stack []*Node
}

func (it *InOrderIterator) HasNext() bool {
	return len(it.stack) > 0
}

func (it *InOrderIterator) Next() *Node {
	if !it.HasNext() {
		return nil
	}
	// Pop top
	n := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]
	// If the node has a right subtree, push its leftmost path
	if n.Right != nil {
		it.pushLeft(n.Right)
	}
	return n
}

// pushLeft pushes node and all its left children onto the stack
func (it *InOrderIterator) pushLeft(n *Node) {
	for n != nil {
		it.stack = append(it.stack, n)
		n = n.Left
	}
}

// Client
func RunBinaryTreeInOrderIterator() {
	// Build a simple tree:
	//      4
	//     / \
	//    2   6
	//   / \ / \
	//  1  3 5  7
	root := &Node{4,
		&Node{2, &Node{1, nil, nil}, &Node{3, nil, nil}},
		&Node{6, &Node{5, nil, nil}, &Node{7, nil, nil}},
	}
	tree := NewBinaryTree(root)
	it := tree.Iterator()
	for it.HasNext() {
		node := it.Next()
		fmt.Print(node.Value, " ")
	}
	// Output: 1 2 3 4 5 6 7
}
