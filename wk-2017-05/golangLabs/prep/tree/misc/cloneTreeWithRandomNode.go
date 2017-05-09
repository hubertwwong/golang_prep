package main

import "fmt"

func createTree01() *Node {
  var t Tree
  t.insertList([]int{3,2,4,1,5})
  t.linkRandom(3,5)
  return t.root
}

// Tree stuff
// ==========================

type Node struct {
  val int
  left *Node
  right *Node
  random *Node
}

type Tree struct {
  root *Node
}

func (t *Tree) insert(val int) {
  if t.root == nil {
    t.root = &Node{val, nil, nil, nil}
  } else {
    for curNode := t.root ; curNode != nil ; {
      if val > (*curNode).val {
        if (*curNode).right != nil {
          curNode = (*curNode).right
        } else {
          (*curNode).right = &Node{val, nil, nil, nil}
          break
        }
      } else {
        if (*curNode).left != nil {
          curNode = (*curNode).left
        } else {
          (*curNode).left = &Node{val, nil, nil, nil}
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

func (t *Tree) find(val int) *Node {
  return t.findDFS(val, t.root)
}

func (t *Tree) linkRandom(src, dest int) {
  srcNode := t.find(src)
  destNode := t.find(dest)

  if srcNode != nil && destNode != nil {
    (*srcNode).random = destNode
  }
}

func (t *Tree) findDFS(val int, root *Node) *Node {
  if root == nil {
    return nil
  } else if val == (*root).val {
    return root
  } else if val > (*root).val {
    return t.findDFS(val, (*root).right)
  } else {
    return t.findDFS(val, (*root).left)
  }
}

// Acting on tree stuff...
// =======================

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

func cloneTree(root *Node) *Node {
  visited := make(map[int]*Node)
  return cloneTreeDFS(root, visited)
}

func cloneTreeDFS(root *Node, visited map[int]*Node) *Node {
  if root == nil {
    return nil
  }

  rootVal := (*root).val
  if val, ok := visited[rootVal]; ok {
    return val
  }

  // copy root node...
  newRoot := Node{rootVal, nil, nil, nil}
  visited[rootVal] = &newRoot

  newRoot.left = cloneTreeDFS((*root).left, visited)
  newRoot.right = cloneTreeDFS((*root).right, visited)
  newRoot.random = cloneTreeDFS((*root).random, visited)

  return &newRoot
}

func main() {
  t1 := createTree01()
  t2 := cloneTree(t1)

  fmt.Println("> Original")
  printBFS(t1)
  
  fmt.Println("> Cloned")
  printBFS(t2)
}

/*

Clone a tree with a random node pointer.
Almost the same as the tree but you need a hash table to refresh the cloned nodes.
This is actually a little different from the clone graph q?
Simpliest way might be to treat it not that differently from the graph and just assume 3.
instead of n.
I think this way will work with b-trees and not just bst...

Assuming a bst for now just to simplify the traversal to find the node for the insert random.

*/