package main

import "fmt"

type Node struct {
	left *Node
	right *Node
	val int
}

// new Node
func newNode(val int) (*Node) {
	n := Node{nil, nil, val}
	return &n
}

// insert method
func insert(n *Node, val int) (*Node) {
	if n == nil {
		return newNode(val)
	}

	// left or right.
	nodeVal := (*n).val
	if val < nodeVal {
		(*n).left = insert((*n).left, val)
	} else {
		(*n).right = insert((*n).right, val)
	}

	// return node.
	return n
}

// insert multile item
func insertM(vals []int) (*Node) {
	var n *Node
	for i:=0 ; i<len(vals) ; i++ {
		n = insert(n, vals[i])
	}
	return n
}

// print tree method
func printTreeBF(n *Node) {
	if n != nil {
		fmt.Println(">", (*n).val)
	}
	printTreeBF((*n).left)
	printTreeBF((*n).right)
}

// helper func to return the max value.
func max(x, y int) (int) {
	if x > y {
		return x
	} else {
		return y
	}
}

// max depth method
func maxDepth(n *Node) (int) {
	if n == nil {
		return 0
	}

	if (*n).left == nil && (*n).right == nil {
		return 1
	}

	
	leftCount := maxDepth((*n).left)
	rightCount := maxDepth((*n).right)
	curCount := max(leftCount, rightCount)

	// backing out. add 1 because you  are counting this node.
	return curCount + 1
}

func main() {
	maxVal := maxDepth(nil)
	fmt.Println(">0", maxVal)

	l1 := []int{42}
	n := insertM(l1)
	maxVal = maxDepth(n)
	fmt.Println(">1", maxVal)

	l2 := []int{2,1,3}
	n = insertM(l2)
	maxVal = maxDepth(n)
	fmt.Println(">2", maxVal)

	l3 := []int{2,1,3,4,5}
	n = insertM(l3)
	maxVal = maxDepth(n)
	fmt.Println(">3", maxVal)
}