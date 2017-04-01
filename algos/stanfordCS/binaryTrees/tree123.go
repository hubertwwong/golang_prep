package main

import "fmt"

type Node struct {
	left *Node
	right *Node
	val int
}

// This returns a node with the value set.
// Just a shortcut.
func newNode(val int) (*Node) {
	newNode := Node{nil, nil, val}
	return &newNode
}

func insert(n *Node, val int) (*Node) {
	if n == nil {
		return newNode(val);
	}

	// set the left or right pointer to something.
	if val < (*n).val {
		(*n).left = insert((*n).left, val)
	} else {
		(*n).right = insert((*n).right, val)
	}

	// node unchanged.
	return n
}

func lookup(n *Node, val int) (bool) {
	if n == nil {
		return false
	}

	if (*n).val == val {
		return true
	}

	if val < (*n).val {
		return lookup((*n).left, val)
	} else {
			return lookup((*n).right, val)
	}
}

func printTreeBF(n *Node) {
	if n != nil {
		fmt.Println(">", (*n).val)
	}
	if (*n).left != nil {
		printTreeBF((*n).left)
	}
	if (*n).right != nil {
		printTreeBF((*n).right)
	} 
}

func main() {
	var root *Node

	// insert 3 items.
	root = insert(root, 2)
	root = insert(root, 1)
	root = insert(root, 3)

	// how to print tree..
	printTreeBF(root)

	// lookup test
	x := lookup(root, 3)
	fmt.Println("3 is ", x)
	x = lookup(root, 20)
	fmt.Println("20 is ", x)
}