package main

import "fmt"

type Node struct {
  p *Node
  v int
}

// This was changed to take pointers.
// You need to modify the actual address and not copies.
func addLink(from, to *Node) (*Node) {
  toAddress := to;
  (*from).p = toAddress;
  return from;
}


// create a list of nodes.
// starts on 1 at the end of the list.
// so it will count backwards.
func createList(numNodes int) (*Node) {
  var startNode, curNode *Node

  // create last linked
  lastNode := Node{nil, numNodes}

  // handle edge cases.
  if numNodes < 0 {
    return nil        
  } else if numNodes == 1 {
    return &lastNode
  }

  // create n-1 other node.
  for i := 0 ; i<numNodes ; i++ {
    //fmt.Println("i", i)
    // handle init case.
    if i == 0 {
      curNode = &lastNode
    }

    // setup new node.
    newNode := Node{nil, numNodes-(i+1)}

    // link node.
    curNode = addLink(&newNode, curNode)

    // update start node
    startNode = curNode
  }

  return startNode
}

// assming its positive for now.
func createListOfNumbers(myNum int) (*Node) {
  var startNode, curNode *Node
  
  // initial cases.
  if myNum < 0 {
    return nil
  } else if myNum < 10 {
    lastNode := Node{nil, myNum}
    curNode = &lastNode
    return curNode
  } else {
    lastNode := Node{nil, myNum%10}
    fmt.Println(">i", myNum%10)
    curNode = &lastNode
  }

  for curNum := myNum/10 ; ; curNum = curNum/10 {
    if curNum >= 10 {
      fmt.Println("l", curNum)
      // setup new node.
      newNode := Node{nil, curNum%10}
      
      // link node.
      curNode = addLink(&newNode, curNode)
      
      // update start node
      startNode = curNode
    } else {
      fmt.Println("f", curNum)
      // setup new node.
      newNode := Node{nil, curNum}
      
      // link node.
      curNode = addLink(&newNode, curNode)

      // update start node
      startNode = curNode

      // number is < 10. break out of loop.
      break
    }
  }

  return startNode
}

func count(head *Node, n int) (int) {
  // initial checks. If node does not exit, return -1.
  if head == nil {
    return -1
  }

  count := 0

  for curPoint := head ;  ; curPoint = (*curPoint).p {
    if (*curPoint).p != nil {
      // normal case.
      if (*curPoint).v == n {
        count++
      }
    } else {
      // end node case.
      if (*curPoint).v == n {
        count++
      }
      // break out of for look. you are on the last node.
      break
    }
  }

  return count
}

func main() {
  // defining a var to set below for current number.
  var curNum int
  num100234555 := createListOfNumbers(100234555)

  fmt.Println("> SOME BASIC CHECKS ON", 100234555)

  curNum = 0
  fmt.Println("num of", curNum, "in the list is", count(num100234555, curNum))

  curNum = 5
  fmt.Println("num of", curNum, "in the list is", count(num100234555, curNum))

  curNum = 3
  fmt.Println("num of", curNum, "in the list is", count(num100234555, curNum))

  curNum = 30
  fmt.Println("num of", curNum, "in the list is", count(num100234555, curNum))

  fmt.Println("> ERROR CASES.")

  fmt.Println("nil should return -1", count(nil, curNum))

  numSingleDig := createListOfNumbers(1)
  curNum = 30
  fmt.Println("num of", curNum, "in the list is", count(numSingleDig, curNum))

  curNum = 1
  fmt.Println("num of", curNum, "in the list is", count(numSingleDig, curNum))
}