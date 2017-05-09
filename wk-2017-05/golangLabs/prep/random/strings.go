package main

import "fmt"

type Node struct {
  val byte
  word bool
}

func main() {
  a := "hello world"
  fmt.Println(string(a[2]))

  // hash...
  h := make(map[byte]int)
  h[a[2]]=42
  fmt.Println(h,h[a[2]])
  fmt.Println(h[108])

  // hash of pointers.
  h2 := make(map[byte]*Node)

  // make a Node
  n1 := Node{43, true}
  h2[a[2]] = &n1
  fmt.Println(h2[a[2]], h2[3])
}