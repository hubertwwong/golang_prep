package main

import "fmt"

// == QUEUE ==

type NodeQ struct {
  next *NodeQ
  val *Node
}

func newNodeQ(val *Node) (*NodeQ) {
  n := NodeQ{nil, val}
  return &n
} 

func enqueue(val *Node, root *NodeQ) (*NodeQ) {
  if val == nil && root == nil {
    return nil
  } else if root == nil {
    n := newNodeQ(val)
    return n
  } else {
    n := newNodeQ(val)
    (*n).next = root
    return n
  }
}

func dequeue(root *NodeQ) (*Node, *NodeQ) {
  if root == nil {
    return nil, nil
  }

  curV := (*root).val
  curN := root
  newRoot := (*root).next
  // remove the reference. I wonder if this does the garbage collection.
  (*curN).next = nil

  return curV, newRoot
}



// == PRINT BFS ==

func printBFS(root *Node) {
  if root != nil {
    var qCurrent *NodeQ
    var qNext *NodeQ
    var nCurrent *Node

    qCurrent = enqueue(root, qCurrent)

    // overall
    for {
      // current row
      for {
        // exit condition
        if qCurrent == nil {
          break
        }

        // current item
        nCurrent, qCurrent = dequeue(qCurrent)

        // print the item.
        fmt.Println("> BST >", (*nCurrent).val)

        // enqueue the children. if not nil
        if (*nCurrent).left != nil {
          qNext = enqueue((*nCurrent).left, qNext)
        }
        if (*nCurrent).right != nil {
          qNext = enqueue((*nCurrent).right, qNext)
        }
      }

      // check the next queue.
      if qNext == nil {
        break
      } else {
        qCurrent = qNext
        qNext = nil
        fmt.Println("")
      }
    }
  }
}



// == TREE ==

type Node struct {
  left *Node
  right *Node
  val int
}

func newNode(val int) (*Node) {
  n := Node{nil, nil, val}
  return &n
}

// stick to root for the node name.
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



// == MAIN FUNC ==

func sameTree(root1 *Node, root2 *Node) (bool) {
  if root1 != nil && root2 != nil {
    var q1Current, q2Current *NodeQ
    var q1Next, q2Next *NodeQ
    var n1Current, n2Current *Node
    var q1Len, q2Len int

    q1Current = enqueue(root1, q1Current)
    q2Current = enqueue(root2, q2Current)
    q1Len = 0
    q2Len = 0

    // overall
    for {
      // current row
      for {
        // exit condition
        if q1Current == nil || q2Current == nil {
          break
        }

        // current item
        n1Current, q1Current = dequeue(q1Current)
        n2Current, q2Current = dequeue(q2Current)

        // Check for equality.
        fmt.Println((*n1Current).val, (*n2Current).val)
        if (*n1Current).val != (*n2Current).val {
          return false
        }

        // print the item.
        //fmt.Println("> BST >", (*nCurrent).val)

        // enqueue the children. if not nil
        if (*n1Current).left != nil {
          q1Next = enqueue((*n1Current).left, q1Next)
          q1Len++
        }
        if (*n1Current).right != nil {
          q1Next = enqueue((*n1Current).right, q1Next)
          q1Len++
        }
        if (*n2Current).left != nil {
          q2Next = enqueue((*n2Current).left, q2Next)
          q2Len++
        }
        if (*n2Current).right != nil {
          q2Next = enqueue((*n2Current).right, q2Next)
          q2Len++
        }
      }

      // length check...
      if q1Len != q2Len {
        return false
      }
      q1Len = 0
      q2Len = 0

      // check the next queue.
      if q1Next == nil {
        break
      } else {
        q1Current = q1Next
        q1Next = nil
      }
      if q2Next == nil {
        break
      } else {
        q2Current = q2Next
        q2Next = nil
      }
    }
  } else {
    return false
  }

  return true
}



// == HELPERS ==

func runTest(vals1 []int, vals2 []int) {
  //fmt.Println("> runTest")
  l1 := insertList(vals1)
  l2 := insertList(vals2)
  
  fmt.Println(sameTree(l1, l2))
  //printBFS(l)
  //fmt.Println("")
}


func main() {
  // nil on right
  //runTest([]int{2,1,3}, nil)
  // nil on left
  //runTest(nil, []int{2,1,3})
  // same
  runTest([]int{10,5,15,3,7}, []int{10,5,15,7,3})
  // diff node values
  runTest([]int{10,5,15,3,7}, []int{10,5,15,20,13}) 
  // missing a value
  runTest([]int{10,5,15,3}, []int{10,5,15,3,7}) 
  

  //runTest([]int{2,1,3})
  //runTest([]int{1,2,3,4,5})
  //runTest([]int{10,5,15,3,7})
}