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

// I don't think this make any assumptions..
// on the order.
func frontBackSplit(headP *Node) (*Node, *Node) {
  // sanity
  if headP == nil {
    return nil, nil
  }

  // edge cases
  if ((*headP).p == nil) {
    // 1 node.
    return headP, nil
  } else if ((*headP).p != nil && (*((*headP).p)).p == nil) {
    // 2 nodes.
    left := headP
    right := (*headP).p

    // set the end p of the left side to nil
    (*left).p = nil

    return left, right
  }

  // other cases
  for slowP, fastP := headP, headP ; ; {
    // traverse pointers
    if fastP != nil && (*fastP).p != nil && (*((*fastP).p)).p != nil {
      // Normal case.
      fastP = (*((*fastP).p)).p
      slowP = (*slowP).p
    } else if fastP != nil && (*fastP).p != nil {
      // even case

      // save the right side pointer.
      rightP := (*slowP).p
      
      // nil the end of the left side pointer.
      (*slowP).p = nil

      return headP, rightP
    } else {
      // odd case

      // save the right side pointer.
      rightP := (*slowP).p
      
      // nil the end of the left side pointer.
      (*slowP).p = nil
      
      return headP, rightP
    }

  }

  return nil, nil
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
    
    if newLeftP == nil {
      tempP, newRightP = pop(newRightP)
    } else if newRightP == nil {
      tempP, newLeftP = pop(newLeftP)
    } else if (*newLeftP).v < (*newRightP).v {
      // left
      tempP, newLeftP = pop(newLeftP)
    } else {
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

func mergeSort(headP *Node) (*Node) {
  //finalP := headP // do you need this?
  if headP == nil {
    return nil
  } else if (*headP).p == nil {
    // Only 1 node.
    return headP
  } else {
    // split the list in half
    leftP, rightP := frontBackSplit(headP)
    leftP = mergeSort(leftP)
    rightP = mergeSort(rightP)
    headP = sortedMerge(leftP, rightP)
  }
  
  return headP
}



func main() {
  fmt.Println("> 321")
  longList1 := createListOfNumbers(321)
  final1 := mergeSort(longList1)
  printList(final1)
  fmt.Println("")

  fmt.Println("> nil")
  //longList2 := createListOfNumbers(321)
  final2 := mergeSort(nil)
  printList(final2)
  fmt.Println("")

  fmt.Println("> 1111111")
  longList3 := createListOfNumbers(1111111)
  final3 := mergeSort(longList3)
  printList(final3)
  fmt.Println("")

  fmt.Println("> 213231123")
  longList4 := createListOfNumbers(213231123)
  final4 := mergeSort(longList4)
  printList(final4)
  fmt.Println("")

  fmt.Println("> 123456")
  longList5 := createListOfNumbers(123456)
  final5 := mergeSort(longList5)
  printList(final5)
  fmt.Println("")

  fmt.Println("> 12345654321")
  longList6 := createListOfNumbers(12345654321)
  final6 := mergeSort(longList6)
  printList(final6)
  fmt.Println("")
}