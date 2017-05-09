package main

import "fmt"

type Node struct {
  val int
  left *Node
  right *Node
}

type Tree struct {
  root *Node
}

func (t *Tree) insert(val int) {
  if t.root == nil {
    t.root = &Node{val, nil, nil}
  } else {
    for curNode := t.root ; curNode != nil ; {
      if val > (*curNode).val {
        if (*curNode).right != nil {
          curNode = (*curNode).right
        } else {
          (*curNode).right = &Node{val, nil, nil}
          break
        }
      } else {
        if (*curNode).left != nil {
          curNode = (*curNode).left
        } else {
          (*curNode).left = &Node{val, nil, nil}
          break
        }
      }
    }
  }
}

func (t *Tree) insertList(vals []int) {
  for i := 0 ; i < len(vals) ; i ++ {
    t.insert(vals[i])
  }
}

func printBFS(root *Node) {
  var q1 []*Node
  var q2 []*Node

  if root != nil {
    q1 = make([]*Node, 0)
    q1 = append(q1, root)
  }

  for ; len(q1) != 0 ; {
    // q1 to q2
    q2 = make([]*Node, 0)
    for i := 0 ; i < len(q1) ; i++ {
      curNode := *q1[i]
      //fmt.Printf("%3d", curNode.val)
      fmt.Println(curNode)
      if curNode.left != nil {
        q2 = append(q2, curNode.left)
      }
      if curNode.right != nil {
        q2 = append(q2, curNode.right)
      }
    }
    fmt.Println("")

    //assign q2 to q1.
    q1 = q2
  }
}

/*
  clone...
  Do a DFS search...
  
  end of the list.
  - return a new copy
  - on every other node.
  - attach left and right and return it.


*/

func cloneTree(root *Node) *Node {
  if root == nil {
    return nil
  }

  // copy the node.
  newNode := Node{(*root).val, nil, nil}

  // attach the children
  newNode.left = cloneTree((*root).left)
  newNode.right = cloneTree((*root).right)

  // return the node.
  return &newNode
}

func main() {
  var t Tree
  var l []int

  //l = []int{2,4,1,3,5}
  l = []int{2,3,1}
  t.insertList(l)
  printBFS(t.root)

  fmt.Println(">>>> cloned")
  clonedRoot := cloneTree(t.root)
  printBFS(clonedRoot)

  //fmt.Println(t.root)
}