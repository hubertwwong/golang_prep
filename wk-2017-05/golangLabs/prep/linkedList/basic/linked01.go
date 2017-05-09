package main

import "fmt"

type Node struct {
  Point *Node
  Value int
}

func main() {
  // foo1 is the first node.
  foo1 := Node{nil,2}
  foo1_ptr := &foo1
  fmt.Println(">foo1")
  fmt.Println(foo1)
  fmt.Println(foo1_ptr)

  // foo2 is the second node.
  // foo2 points to foo1
  fmt.Println(">foo2")
  foo2 := Node{foo1_ptr,3}
  fmt.Println(foo2);

  // deference first node points to 2.
  out1 := foo2.Point
  fmt.Println(">foo2 point to foo1")
  fmt.Println(out1)
}