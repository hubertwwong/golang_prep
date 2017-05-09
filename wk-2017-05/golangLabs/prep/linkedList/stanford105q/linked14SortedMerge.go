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

func sortedMerge(leftP, rightP *Node) (*Node) {
  // sanity checks
  if leftP == nil && rightP == nil {
    return nil
  } else if leftP == nil && rightP != nil {
    return rightP
  } else if rightP == nil && leftP != nil {
    return leftP
  }

  // do you need this?
  newLeftP := leftP
  newRightP := rightP
  var finalP, tempP, curP *Node
  finalP = nil
  
  // bootstrap the first node in the final list.
  if (*newLeftP).v < (*newRightP).v {
    tempP, newLeftP = pop(newLeftP)
    finalP = tempP
    curP = finalP
  } else {
    tempP, newRightP = pop(newRightP)
    finalP = tempP
    curP = finalP
  }

  // pop is find.
  // push should be append
  for ; newLeftP != nil || newRightP != nil ; {
    tempP = nil // do you need this?
    //fmt.Println(">>>>>>>>>>")
    //fmt.Println("> finalP")
    //printList(finalP)
    //fmt.Println("> newLeftP")
    //printList(newLeftP)
    //fmt.Println("> newRightP")
    //printList(newRightP)
    //fmt.Println("")

    if newLeftP == nil {
      //fmt.Println(">1")
      tempP, newRightP = pop(newRightP)
    } else if newRightP == nil {
      //fmt.Println(">2")
      tempP, newLeftP = pop(newLeftP)
    } else if (*newLeftP).v < (*newRightP).v {
      // fmt.Println(">3")
      // left
      tempP, newLeftP = pop(newLeftP)
    } else {
      //fmt.Println(">4")
      // right
      tempP, newRightP = pop(newRightP)
    }

    // append the node.
    (*curP).p = tempP
    // move curret pointer to end of new list.
    curP = (*curP).p
  }

  return finalP
}



func main() {
  fmt.Println(">>> generic case")
  left1 := createListOfNumbers(123)
  right1 := createListOfNumbers(456)
  final1 := sortedMerge(left1, right1)
  printList(final1)
  fmt.Println("")

  fmt.Println(">>> lefts")
  left2 := createListOfNumbers(123)
  //right2 := createListOfNumbers(456)
  final2 := sortedMerge(left2, nil)
  printList(final2)
  fmt.Println("")

  fmt.Println(">>> rights")
  //left3 := createListOfNumbers(123)
  right3 := createListOfNumbers(456)
  final3 := sortedMerge(nil, right3)
  printList(final3)
  fmt.Println("")

  fmt.Println(">>> some case with repeats")
  left4 := createListOfNumbers(1237777)
  right4 := createListOfNumbers(456)
  final4 := sortedMerge(left4, right4)
  printList(final4)
  fmt.Println("")

  fmt.Println(">>> nil")
  //left5 := createListOfNumbers(123)
  //right5 := createListOfNumbers(456)
  final5 := sortedMerge(nil, nil)
  printList(final5)
  fmt.Println("")

  fmt.Println(">>> some case with same number on both list")
  left6 := createListOfNumbers(1237777)
  right6 := createListOfNumbers(1456)
  final6 := sortedMerge(left6, right6)
  printList(final6)
  fmt.Println("")

}