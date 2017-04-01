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

func insertList(valList []int) (*Node) {
  var root *Node
  root = nil

  for i:=0 ; i<len(valList) ; i++ {
    root = insert(valList[i], root)
  }

  return root
}

// i'm assming -1 is false. for the return
func hasSumPath(sum int, root *Node) (bool) {
  //fmt.Println(">hs", sum, root)
  if root == nil {
    if sum == 0 {
      //fmt.Println(">>>")
      return true
    } else {
      return false
    }
  }

  leftBool := hasSumPath((sum - (*root).val), (*root).left)
  rightBool := hasSumPath((sum - (*root).val), (*root).right)
  return (leftBool || rightBool)
}

// wrapping stuff in test cases.
func testRun(sum int, vals []int) {
  root := insertList(vals)
  result := hasSumPath(sum, root)
  fmt.Println(">", sum, vals, result)
}



func main() {
  testRun(5, []int{2,1,3})
  testRun(3, []int{2,1,3})
  testRun(4, []int{2,1,3})
  testRun(0, []int{2,1,3})
}