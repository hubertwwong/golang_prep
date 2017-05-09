package main

import "fmt"

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
  n := newNode(val)
  (*n).next = root
  return n
}

func insertList(vals []int) (*Node) {
  if len(vals) == 0 {
    return nil
  }
  var root *Node
  for i:=0 ; i<len(vals) ; i++ {
    root = insert(vals[i], root)
  }
  return root
}

func printList(root *Node) {
  for curP := root ; curP != nil ; curP = (*curP).next {
    fmt.Println((*curP).val)
  }
}

func reverse(root *Node) (*Node) {
  if root == nil {
    return nil
  }

  if (*root).next == nil {
    return root
  }

  // we know 2 nodes exist.
  prevP := root
  curP := (*root).next
  nextP := (*curP).next

  // handle 2 node case.
  if nextP == nil {
    (*curP).next = prevP
    (*prevP).next = nil
    return curP
  }

  // more than 3

  
  for {
    fmt.Println(prevP, curP, nextP)
    // this is wrong
    if nextP == nil {
      //(*prevP).next = nil
      (*curP).next = prevP
      return curP
    } else {
      // set the end new node to nil.
      if prevP == root {
        (*prevP).next = nil
      }
      (*curP).next = prevP
      prevP = curP
      curP = nextP
      nextP = (*nextP).next
    }
  }
}

func main() {
  i := insertList([]int{1})
  fmt.Println(">before")
  printList(i)
  j := reverse(i)
  fmt.Println(">after")
  printList(j)
}


/*

1st attemp

1,2,3,4
prints 4...
attempt 2.
no go....
3rd attempt. printing...


*/