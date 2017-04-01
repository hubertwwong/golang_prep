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
    (*root).left = insert(val, (*root).left)
  } else {
    (*root).right = insert(val, (*root).right)
  }

  return root
}

func insertList(vals []int) (*Node) {
  var root *Node
  root = nil

  for i:=0 ; i<len(vals) ; i++ {
    root = insert(vals[i], root)
  }

  return root
}

func printInOrder(root *Node) {
  if root != nil {
    printInOrder((*root).left)
    fmt.Println(">PIO", (*root).val)
    printInOrder((*root).right)
  }
}

func printPostOrder(root *Node) {
  if root != nil {
    printPostOrder((*root).left)
    printPostOrder((*root).right)
    fmt.Println(">PPO", (*root).val)
  }
}

func runner(vals []int) {
  il := insertList(vals)
  fmt.Println(">runner")
  fmt.Println("")
  printInOrder(il)
  fmt.Println("")
  printPostOrder(il)
  fmt.Println("")
}


func main() {
  // l1 := []int{3,2,1}
  // i1 := insertList(l1)
  // printInOrder(i1)

  // l2 := []int{300,1,20,70,30}
  // i2 := insertList(l2)
  // printInOrder(i2)

  // l3 := []int{1,2,3}
  // i3 := insertList(l3)
  // printInOrder(i3)

  // l4 := []int{42}
  // i4 := insertList(l4)
  // printInOrder(i4)

  runner([]int{3,2,1})
  runner([]int{3,20,1})
  runner([]int{4,2,5,1,3})
}





/*

07:08 start
007:27 end
5 mins for connections issue
and the thing seems to work....
runner([]int{3,20,1})

*/