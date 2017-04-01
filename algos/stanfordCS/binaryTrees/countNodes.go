package main

import "fmt"

// struct for node
type Node struct {
	left *Node
	right *Node
	val int
}

// helper to make a node.
func newNode(val int) (*Node) {
	n := Node{nil, nil, val}
	return &n
}

// insert for BST
func insert(n *Node, val int) (*Node) {
	// you are at the leaf node.
	if n == nil {
		return newNode(val)
	}

	// determine if you want to traverse left or right.
	fmt.Println(val, (*n).val)
	if val < (*n).val {
		(*n).left = insert((*n).left, val)
	} else {
		(*n).right = insert((*n).right, val)
	}

	// node was unchanged?
	return n
}

func printTreeBF(n *Node) {
	if n != nil {
		fmt.Println(">>", (*n).val)
	}
	if (*n).left != nil {
		printTreeBF((*n).left)
	}
	if (*n).right != nil {
		printTreeBF((*n).right)
	} 
}

// count nodes function.
func countNodes(n *Node) (int) {
	if n == nil {
		return 0
	}
	
	// defining the pointers to make it easier to read.
	leftP := (*n).left
	rightP := (*n).right
	curCount := 0
	// lazy init for int

	// you hit a leaf node, return 1.
	if leftP == nil && rightP == nil {
		return 1
	}

	// recursion. Go left and right. They should return a int, return the results.
	curCount += countNodes(leftP)
	curCount += countNodes(rightP)

	// return count of left and right nodes.
	// you need to add 1 for this node.
	return curCount + 1
}

func main() {
	curCount := -3

	// 0 nodes.
	curCount = countNodes(nil)
	fmt.Println("> 0", curCount)

	// 1 node.
	root1 := insert(nil, 1)
	curCount = countNodes(root1)
	printTreeBF(root1)
	fmt.Println("> 1", curCount)

	// 3 nodes
	root2 := insert(nil, 4)
	root2 = insert(root2, 3)
	root2 = insert(root2, 5)
	curCount = countNodes(root2)
	printTreeBF(root2)
	//fmt.Println("left pointer in main", (*root2).left)
	fmt.Println("> 3", curCount)
	
	// 5 nodes.
	root3 := insert(nil, 4)
	root3 = insert(root3, 3)
	root3 = insert(root3, 5)
	root3 = insert(root3, 1)
	root3 = insert(root3, 6)
	curCount = countNodes(root3)
	fmt.Println("> 5", curCount)
}