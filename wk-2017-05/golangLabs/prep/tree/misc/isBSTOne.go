package main

import "fmt"

// == tree setup

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
  valsLen := len(vals)
  
  if valsLen == 0 {
    return nil
  }

  var root *Node
  for i := 0 ; i < valsLen ; i++ {
    root = insert(vals[i], root)
  }
  return root
}

// == helper

func minNodeVal(root *Node) (int) {
  if root == nil {
    return -1
  }
  // child node.
  // this might not let you get back...
  if (*root).left == nil && (*root).right == nil {
    return (*root).val
  }

  // dfs
  leftV := minNodeVal((*root).left)
  rightV := minNodeVal((*root).right)

  // both are nil
  if leftV == -1 && rightV == -1 {
    return (*root).val 
  } else if leftV == -1 {
    if rightV < (*root).val {
      return rightV
    } else {
      return (*root).val
    }
  } else if rightV == -1 {
    if leftV < (*root).val {
      return leftV
    } else {
      return (*root).val
    }
  } else if leftV < rightV {
    if leftV < (*root).val {
      return leftV
    } else {
      return (*root).val
    }
  } else {
    if rightV < (*root).val {
      return rightV
    } else {
      return (*root).val
    }
  }
}

func maxNodeVal(root *Node) (int) {
  if root == nil {
    return -1
  }
  // child node.
  // this might not let you get back...
  if (*root).left == nil && (*root).right == nil {
    return (*root).val
  }

  // dfs
  leftV := maxNodeVal((*root).left)
  rightV := maxNodeVal((*root).right)

  // both are nil
  if leftV == -1 && rightV == -1 {
    return (*root).val 
  } else if leftV == -1 {
    if rightV > (*root).val {
      return rightV
    } else {
      return (*root).val
    }
  } else if rightV == -1 {
    if leftV > (*root).val {
      return leftV
    } else {
      return (*root).val
    }
  } else if leftV > rightV {
    if leftV > (*root).val {
      return leftV
    } else {
      return (*root).val
    }
  } else {
    if rightV > (*root).val {
      return rightV
    } else {
      return (*root).val
    }
  }
}

// == main

func isBST(root *Node) (bool) {
  if root == nil {
    return true
  }
  // this will ensure that all children return true
  if (*root).left == nil && (*root).right == nil  {
    return true
  }

  leftBST := isBST((*root).left)
  rightBST := isBST((*root).right)

  //fmt.Println("> isbst >")
  //fmt.Println(root)
  //fmt.Println(leftBST)
  //fmt.Println(rightBST)

  if leftBST && rightBST {
    leftV := minNodeVal((*root).left)
    rightV := maxNodeVal((*root).right)
    //fmt.Println(leftV, rightV) 

    // main check. not that min and max values return -1 child nodes.
    if leftV != -1 && rightV != -1 { 
      if leftV < (*root).val && (*root).val < rightV {
        return true
      } else {
        return false
      }
    } else if leftV != -1 {
      if leftV < (*root).val {
        return true
      } else {
        return false
      }
    } else if rightV != -1 {
      if (*root).val < rightV {
        return true
      } else {
        return false
      }
    } else {
      // both are -1.
      return true
    }
  } else {
    // if you see a false value, just return it.
    return false
  }
}

func testRun(vals []int) {
  tree1 := insertList(vals)
  fmt.Println(minNodeVal(tree1), maxNodeVal(tree1))
  fmt.Println(isBST(tree1))
}

func main() {
  //testRun([]int{42,2,1,3,32})
  //testRun([]int{2})
  //testRun([]int{2,1,3})
  //testRun([]int{3,2,1})

  // broken tree test.
  // and this does work.
  // root := newNode(2)
  // left := newNode(1)
  // right := newNode(3)
  // (*root).right = left
  // (*root).left = right
  // fmt.Println(isBST(root))

  root2 := insertList([]int{4,3,2})
  // change this to a 1 and it will be false.
  right2 := newNode(6)
  (*root2).right = right2
  fmt.Println(isBST(root2))
}