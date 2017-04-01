package main

import "fmt"

type Node struct {
	left *Node
	right *Node
	val int
}

func newNode(val int) (*Node) {
	n := Node{nil, nil, val}
	return &n
}

func insert(val int, root *Node) (*Node) {
	if root == nil {
		return newNode(val)
	}
	
	if val <= (*root).val {
		fmt.Println("left")
		(*root).left = insert(val, (*root).left)
	} else {
		fmt.Println("right")
		(*root).right = insert(val, (*root).right)
	}

	return root
}

func inserts(vals []int) (*Node) {
	var root *Node
	
	for i:=0 ; i<len(vals) ; i++ {
		root = insert(vals[i], root)
	}
	
	return root 
}

func mirror(root *Node) {
	if root != nil {
		tmp := (*root).left
		(*root).left = (*root).right
		(*root).right = tmp
		
		mirror((*root).right)
		mirror((*root).left)
	}
}

// this is wrong...
// didn't realize this...
func printBST(root *Node) {
	if root != nil {
		fmt.Println((*root).val)
		printBST((*root).left)
		printBST((*root).right)
	}
}

func testRun(vals []int) {
	fmt.Println("> start")
	root := inserts(vals)
	// quick debug
	fmt.Println((*root).val)
	fmt.Println((*root).left)
	fmt.Println((*root).right)
	printBST(root)
	mirror(root)
	fmt.Println("> end")
	printBST(root)
}

func main() {
	//testRun([]int{2,1,3})
	//testRun([]int{1,2,3})
	testRun([]int{10,5,15,2})
	//testRun([]int{10,5,15,2,7,20,13})	
}


/*

11:52a start
+1 for a gaming switch video
12:03a
12:14 done.

2 small typo erros
but it seems to work...

*/