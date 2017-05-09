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

// this is wrong..
// you don't need a return.
func removeDupe(root *Node) {

  if root != nil {
    // fuck... should bail one early..
    for curP := root ; curP != nil ; curP = (*curP).next {
      fmt.Println("curP", curP)
      // inner loop should not assign to outer loop...
      for prevP, nextP := curP, (*curP).next ; nextP != nil ; {
        fmt.Println(">NP PP",nextP, prevP)
        if (*nextP).val == (*curP).val {
          fmt.Println("deleting")
          tempP := nextP
          nextP = (*nextP).next
          (*tempP).next = nil
          (*prevP).next = nextP
        } else {
          // this increment should only happen if you are not deleting something...
          fmt.Println("increment")
          prevP = (*prevP).next
          nextP = (*nextP).next
        }
      }
      fmt.Println(">>>>>>>")
      printList(root)
    }
  }
}

func printList(root *Node) {
  for curP := root ; curP != nil ; curP = (*curP).next {
    fmt.Println((*curP).val)
  }
}

func main() {
  list1 := insertList([]int{4,4,3,2,1,4,4})
  fmt.Println(">before")
  printList(list1)
  removeDupe(list1)
  fmt.Println(">after")
  printList(list1)
}