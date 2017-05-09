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

func sortedIntersect(leftP, rightP *Node) (*Node) {
  // sanity checks
  if leftP == nil || rightP == nil {
    return nil
  }

  // new list.
  var finalP, curP *Node
  finalP = nil
  curP = nil
  prevInsNum := -1;
  // you need to keep track of the curIntersectNum to avoid duplicates.

  for ; leftP != nil || rightP != nil ; {
    if leftP == nil || rightP == nil {
      // if one of the pointers are nil, you have nothing to compare.
      // break out of the loop and return anything if exist.
      break
    } else if (*leftP).v < (*rightP).v {
      leftP = (*leftP).p
    } else if (*leftP).v > (*rightP).v {
      rightP = (*rightP).p
    } else {
      // creating a new node to insert.
      // does not matter what side you are getting the value from since they are both the same.
      newNode := Node{nil, (*leftP).v}
      tempP := &newNode

      // this top check if for duplicate values.
      if prevInsNum != (*leftP).v {
        if finalP == nil {
          // nothing in the list yet.
          curP = tempP
          finalP = tempP
        } else {
          // assign the new node.
          (*curP).p = tempP
          // traverse the current pointer to the end of the list.
          curP = (*curP).p
        }
      }

      // update the previously inserted number.
      prevInsNum = (*leftP).v

      // this is where both are equal
      leftP = (*leftP).p
      rightP = (*rightP).p
    }
  }

  // return final list
  return finalP
}



func main() {
  leftN1 := 123
  rightN1 := 123
  fmt.Println(">", leftN1, ">", rightN1)
  leftL1 := createListOfNumbers(leftN1)
  rightL1 := createListOfNumbers(rightN1)
  final1 := sortedIntersect(leftL1, rightL1)
  printList(final1)
  fmt.Println("")

  leftN2 := 123
  rightN2 := 456
  fmt.Println(">", leftN2, ">", rightN2)
  leftL2 := createListOfNumbers(leftN2)
  rightL2 := createListOfNumbers(rightN2)
  final2 := sortedIntersect(leftL2, rightL2)
  printList(final2)
  fmt.Println("")

  leftN3 := 123
  //rightN3 := 123
  fmt.Println(">left only")
  leftL3 := createListOfNumbers(leftN3)
  //rightL3 := createListOfNumbers(rightN3)
  final3 := sortedIntersect(leftL3, nil)
  printList(final3)
  fmt.Println("")

  //leftN4 := 123
  rightN4 := 123
  fmt.Println(">right only")
  //leftL4 := createListOfNumbers(leftN4)
  rightL4 := createListOfNumbers(rightN4)
  final4 := sortedIntersect(nil, rightL4)
  printList(final4)
  fmt.Println("")


  //leftN5 := 123
  //rightN5 := 123
  fmt.Println(">double nil")
  //leftL5 := createListOfNumbers(leftN5)
  //rightL5 := createListOfNumbers(rightN5)
  final5 := sortedIntersect(nil, nil)
  printList(final5)
  fmt.Println("")

  leftN6 := 123
  rightN6 := 345
  fmt.Println(">", leftN6, ">", rightN6)
  leftL6 := createListOfNumbers(leftN6)
  rightL6 := createListOfNumbers(rightN6)
  final6 := sortedIntersect(leftL6, rightL6)
  printList(final6)
  fmt.Println("")

  leftN7 := 123779
  rightN7 := 345779
  fmt.Println(">", leftN7, ">", rightN7)
  leftL7 := createListOfNumbers(leftN7)
  rightL7 := createListOfNumbers(rightN7)
  final7 := sortedIntersect(leftL7, rightL7)
  printList(final7)
  fmt.Println("")
}