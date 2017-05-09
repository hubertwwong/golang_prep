package main

import "fmt"

type Node struct {
  p *Node
  v int
}

// Points 'from' -> 'to' node and return the 'from' node.
func addLink(from, to Node) (Node) {
  toLoc := &to;
  from.p = toLoc;
  return from;
}

// Given an initialNode, print the linked list.
func printList(initNode Node) {
  curNode := initNode

  for {
    if curNode.p == nil {
      fmt.Println(curNode.v)
      break
    }

    fmt.Println(curNode.v)
    curNode = *curNode.p
  }
}

// create a list of nodes.
// starts on 1 at the end of the list.
// so it will count backwards.
func createList(numNodes int) (Node) {
  var startNode, curNode Node

  // create last linked
  lastNode := Node{nil, 1}

  // create n-1 other node.
  for i := 0 ; i<numNodes-1 ; i++ {
    // handle init case.
    if i == 0 {
      curNode = lastNode
    }

    // setup new node.
    newNode := Node{nil, curNode.v + 1}

    // link node.
    curNode = addLink(newNode, curNode)

    // update start node
    startNode = curNode
  }

  return startNode
}

func main() {
  // foo1 := Node{nil, 2}
  // foo2 := addLink(Node{nil, 3}, foo1)
  // foo3 := addLink(Node{nil, 4}, foo2)

  fooList := createList(10)

  printList(fooList)
}

/*
You can create a linked list and print it..
Next is to create a circular link.

*/