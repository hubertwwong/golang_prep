package main

import "fmt"

// Linked list struct

type Node struct {
  next *Node
  val int
}

func newNode(val int) (*Node) {
  n := Node{nil, val}
  return &n
}

func insert(val int, root *Node) (*Node) {
  if root == nil {
    return newNode(val)
  }

  //if val <= (*root).val {
  //  (*root).left = insert(val, (*root).left)
  //} else {
  //  (*root).right = insert(val, (*root).right)
  //}

  for {
    if (*root).next != nil {
      root = (*root).next
    } else {
      (*root).next = newNode(val)
      break
    }
  }

  return root
}

func insertList(vals []int) (*Node) {
  valsLen := len(vals)
  if valsLen <= 0 {
    return nil
  }

  var root *Node
  for i:=0 ; i<valsLen ; i++ {
    root = insert(vals[i], root)
  }
  return root
} 

// Main func
func secondToLast(root *Node) (int) {
  if root == nil {
    // you pass it a nil for some reason.
    return -1
  } else if (*root).next == nil {
    // you pass it a single node.
    return -1
  }

  // at this point you can assume you have 2 nodes.
  curN := (*root).next
  prevN := root
  for {
    if (*curN).next != nil {
      curN = (*curN).next
      prevN = (*prevN).next
    } else {
      return (*prevN).val
    }
  }
}

func testRun(vals []int) {
  i := insertList(vals)
  j := secondToLast(i)
  fmt.Println(j)
}

func main() {
  testRun(nil)
  testRun([]int{1})
  testRun([]int{1,2})
  testRun([]int{1,2,3})
  testRun([]int{1,2,3,4})
}



/*

print second to last node...
12:08a start
12:22 lunch....

01:37 back
01:43 debugging...
put up a tree for a linnked list probme....

*/