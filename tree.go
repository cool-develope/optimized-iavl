package iavl

import "fmt"

type Node struct {
	left   *Node
	right  *Node
	parent *Node
	key    uint
	path   uint
	height uint
}

type Tree struct {
	root *Node
	size uint
}

func NewTree() *Tree {
	return &Tree{
		root: nil,
		size: 0,
	}
}

func (t *Tree) NewNode(key uint) *Node {
	t.size++
	return &Node{
		key:  key,
		path: t.size,
	}
}

func (t *Tree) AddNode(key uint) {
	if t.root == nil {
		t.root = t.NewNode(key)
		return
	}
	t.addNode(t.root, t.NewNode(key))
	if t.root.parent != nil {
		t.root = t.root.parent
	}
}

func (t *Tree) addNode(root *Node, node *Node) {
	// is Leaf
	if root.height == 0 {
		if node.key > root.key {
			root.right = node
			root.left = t.NewNode(root.key)
		} else {
			root.left = node
			root.right = t.NewNode(root.key)
			root.key = node.key
		}
		root.height = 1
		root.left.parent = root
		root.right.parent = root
	} else {
		if node.key > root.key {
			t.addNode(root.right, node)
		} else {
			t.addNode(root.left, node)
		}
		lh, rh := root.left.height, root.right.height
		if lh > 1+rh {
			if root.left.right.height > root.left.left.height {
				rightRotate(root.left, root.left.right)
			}
			leftRoate(root, root.left)
		} else if rh > 1+lh {
			if root.right.left.height > root.right.right.height {
				leftRoate(root.right, root.right.left)
			}
			rightRotate(root, root.right)
		} else {
			updateHeight(root)
		}
	}
}

func updateHeight(node *Node) {
	node.height = node.left.height + 1
	if node.right.height > node.left.height {
		node.height = node.right.height + 1
	}
}

func swapParent(parent, child *Node) {
	p := parent.parent
	if p != nil {
		if p.left == parent {
			p.left = child
		} else {
			p.right = child
		}
	}
	child.parent = p
	parent.parent = child
}

func leftRoate(parent, left *Node) {
	parent.path, left.path = left.path, parent.path
	parent.left = left.right
	left.right.parent = parent
	left.right = parent
	swapParent(parent, left)
	updateHeight(parent)
	updateHeight(left)
}

func rightRotate(parent, right *Node) {
	parent.path, right.path = right.path, parent.path
	parent.right = right.left
	right.left.parent = parent
	right.left = parent
	swapParent(parent, right)
	updateHeight(parent)
	updateHeight(right)
}

func (n *Node) print() {
	fmt.Printf("Node %d: height %d path %d", n.key, n.height, n.path)
}

func print(root *Node) {
	if root == nil {
		return
	}
	root.print()
	if root.parent != nil {
		fmt.Print(" Parent ")
		root.parent.print()
	}
	if root.left != nil {
		fmt.Print(" Left ")
		root.left.print()
		fmt.Print(" Right ")
		root.right.print()
	}
	fmt.Println()
	print(root.left)
	print(root.right)
}
