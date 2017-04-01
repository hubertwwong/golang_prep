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
  //fmt.Println(">insert", val, root)
  if root == nil {
    return newNode(val)
  }

  // left right traversal
  if val <= (*root).val {
    fmt.Println(">insert left")
    (*root).left = insert(val, (*root).left)
  } else {
    fmt.Println(">insert right")
    (*root).right = insert(val, (*root).right)
  }

  // return itself.
  return root
}

// probably iff on slices
func inserts(vals []int) (*Node) {
  var root *Node
  root = nil
  
  for i:=0 ; i<len(vals) ; i++ {
    root = insert(vals[i], root)
  }

  //fmt.Println("> inserts",root)

  return root
}

// assuming its positive.
func minValue(root *Node) (int) {
  //fmt.Println(">", root)
  if root == nil {
    return -1
  }

  if (*root).left != nil {
    return minValue((*root).left)
  }

  return (*root).val
}



func main() {
  v1 := []int{1,2,3}
  r1 := inserts(v1)
  result1 := minValue(r1)
  fmt.Println("> result1",result1)

  v2 := []int{3,2,1}
  r2 := inserts(v2)
  result2 := minValue(r2)
  fmt.Println("> result2", result2)

  v3 := []int{2,1,3}
  r3 := inserts(v3)
  result3 := minValue(r3)
  fmt.Println("> result3", result3)

  v4 := []int{9,3,7,2,5,8,4}
  r4 := inserts(v4)
  result4 := minValue(r4)
  fmt.Println("> result4", result4)

  v5 := []int{9}
  r5 := inserts(v5)
  result5 := minValue(r5)
  fmt.Println("> result5", result5)

}


/*
  06:15p
  06:36
  4 errors
  2 passing in wrong args to a func that i created.
  1. typo.
  1. typo in passing in a func that did not exist.

  06:40 debug
  06:53 done.
  got the insert wrong.
  the intial check chould be for the node direct value instead of the child.
  value needs to be larger smaller than node value and then you go left.
  otherwise go right.

  07:02done
  not too bad.
  have not done this in a week or so...
    
*/