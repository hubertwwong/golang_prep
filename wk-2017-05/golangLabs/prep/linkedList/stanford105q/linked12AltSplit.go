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

func altSplit(headP *Node) (*Node, *Node) {
  // sanity check.
  if headP == nil {
    return nil, nil
  }

  var leftP, rightP *Node
  oddEven := 0
  leftP = nil
  rightP = nil

  for curP := headP ; ; {
    // push on one list or the other.
    if oddEven%2 == 0  {
      leftP, curP = moveNode(leftP, curP)
    } else {
      rightP, curP = moveNode(rightP, curP)
    }

    // exit check.
    if curP == nil {
      return leftP, rightP
    }

    // increment.
    oddEven++
  }
}


func main() {
  fmt.Println(">>> generic case")
  list1 := createListOfNumbers(12345)
  finalLeft1, finalRight1 := altSplit(list1)
  fmt.Println("> left")
  printList(finalLeft1)
  fmt.Println("> right")
  printList(finalRight1)
  fmt.Println("")

  fmt.Println(">>> 2 item")
  list2 := createListOfNumbers(12)
  finalLeft2, finalRight2 := altSplit(list2)
  fmt.Println("> left")
  printList(finalLeft2)
  fmt.Println("> right")
  printList(finalRight2)
  fmt.Println("")

  fmt.Println(">>> 1 item")
  list3 := createListOfNumbers(1)
  finalLeft3, finalRight3 := altSplit(list3)
  fmt.Println("> left")
  printList(finalLeft3)
  fmt.Println("> right")
  printList(finalRight3)
  fmt.Println("")

  fmt.Println(">>> nil case")
  //list1 := createListOfNumbers(123456)
  finalLeft4, finalRight4 := altSplit(nil)
  fmt.Println("> left")
  printList(finalLeft4)
  fmt.Println("> right")
  printList(finalRight4)
  fmt.Println("")
}