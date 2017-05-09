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

func getNth(head *Node, num int) (*Node) {
  // preconditions
  if head == nil {
    return nil
  } else if num < 0 {
    return nil
  }

  // init vars
  curIndex := 1

  for curP := head ; ; curP = (*curP).p {
    // check for value.
    // basically you want this checck to happen before the next pointer check.
    if curIndex == num {
      return curP
    }
    
    // Last node.
    if (*curP).p == nil {
      break
    }

    // index increment
    curIndex++
  }

  // catch all. You went thru the list.
  return nil
}



func main() {
  var nthNode int

  num1To9 := createListOfNumbers(34567)

  // common case
  fmt.Println("> Common number. Multiple node.")
  nthNode = 3
  fmt.Println("node ", nthNode, "is", getNth(num1To9, nthNode))
  nthNode = 100
  fmt.Println("node ", nthNode, "is", getNth(num1To9, nthNode))
  nthNode = -3
  fmt.Println("node ", nthNode, "is", getNth(num1To9, nthNode))

  // 1 node case
  fmt.Println("> 1 digit")
  num1 := createListOfNumbers(3)
  nthNode = 1
  fmt.Println("node ", nthNode, "is", getNth(num1, nthNode))
  nthNode = 2
  fmt.Println("node ", nthNode, "is", getNth(num1, nthNode))


  // error checks
  fmt.Println("> error check")
  nthNode = 3
  fmt.Println("nil node ", nthNode, "is", getNth(nil, nthNode))
}