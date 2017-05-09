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
    //fmt.Println(">i", myNum%10)
    curNode = &lastNode
  }

  for curNum := myNum/10 ; ; curNum = curNum/10 {
    if curNum >= 10 {
      //fmt.Println("l", curNum)
      // setup new node.
      newNode := Node{nil, curNum%10}
      
      // link node.
      curNode = addLink(&newNode, curNode)
      
      // update start node
      startNode = curNode
    } else {
      //fmt.Println("f", curNum)
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

func printList(initNode *Node) {
  curNode := initNode
  
  // initial checks
  if initNode == nil {
    fmt.Println("nil")
  } else {
    for {
      // print last node.
      if curNode.p == nil {
        fmt.Println("{nil",(*curNode).v, "}")
        break
      }

      // print anything else
      fmt.Println(*curNode)
      curNode = (*curNode).p
    }
  }
}

// fixing this to return the do both.
// 1st return the top node.
// 2nd return the rest of the list.
func pop(headP *Node) (*Node, *Node){
  // sanity checks
  if headP == nil {
    return nil, nil
  }

  // 1 node.
  if (*headP).p == nil {
    return headP, nil
  }

  // general case..
  newHeadP := (*headP).p
  poppedP := headP
  (*poppedP).p = nil 
  
  return poppedP, newHeadP
}

// it should return the new head node.
func push(headP, newP *Node) (*Node){
  // sanity checks
  if headP == nil || newP == nil {
    return nil
  }

  (*newP).p = headP

  return newP
}

// moves 1 node formP list to the toP list.
// returns the headP of both list. 
func moveNode(toP, fromP *Node) (*Node, *Node) {
  // sanity
  if fromP == nil {
    // You can't move anything.
    // We are making an assumption that both items exist.
    return toP, fromP
  }

  // new pointers.
  var newFromP, newToP *Node

  if toP == nil {
    // Nothing in toP list
    return pop(fromP)
  } else if (*fromP).p == nil {
    // only 1 node in from node.
    newToP = push(toP, fromP)
    return newToP, nil
  } else {
    // generic case. from has multiple nodes.
    // to has at least one node.
    var tempP *Node
    tempP, newFromP = pop(fromP)
    newToP = push(toP, tempP)
    return newToP, newFromP
  }
}



func main() {
  fmt.Println("> generic case.")
  to1 := createListOfNumbers(123)
  from1 := createListOfNumbers(456)
  finalTo1, finalFrom1 := moveNode(to1, from1)
  fmt.Println("> from")
  printList(finalFrom1)
  fmt.Println("> to")
  printList(finalTo1)
  fmt.Println("")

  fmt.Println("> no to.")
  //to2 := createListOfNumbers(123)
  from2 := createListOfNumbers(123)
  finalTo2, finalFrom2 := moveNode(nil, from2)
  fmt.Println("> from")
  printList(finalFrom2)
  fmt.Println("> to")
  printList(finalTo2)
  fmt.Println("")


  fmt.Println("> no from. nothing should change.")
  to3 := createListOfNumbers(123)
  //from3 := createListOfNumbers(123)
  finalTo3, finalFrom3 := moveNode(to3, nil)
  fmt.Println("> from")
  printList(finalFrom3)
  fmt.Println("> to")
  printList(finalTo3)
  fmt.Println("")

  fmt.Println("> 1 item in from.")
  to4 := createListOfNumbers(123)
  from4 := createListOfNumbers(4)
  finalTo4, finalFrom4 := moveNode(to4, from4)
  fmt.Println("> from")
  printList(finalFrom4)
  fmt.Println("> to")
  printList(finalTo4)
  fmt.Println("")
}
