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

// assumptions
// list is sorted ascending
// if its odd, the first return will be larger.
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

      fmt.Println("even case")
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



func main() {
  fmt.Println("> 1 node")
  longList1 := createListOfNumbers(1)
  left1, right1 := frontBackSplit(longList1)
  printList(left1)
  fmt.Println(">>>")
  printList(right1)
  fmt.Println("")

  fmt.Println("> 2 node")
  longList2 := createListOfNumbers(12)
  left2, right2 := frontBackSplit(longList2)
  printList(left2)
  fmt.Println(">>>")
  printList(right2)
  fmt.Println("")

  fmt.Println("> 3 node")
  longList3 := createListOfNumbers(123)
  left3, right3 := frontBackSplit(longList3)
  printList(left3)
  fmt.Println(">>>")
  printList(right3)
  fmt.Println("")

  fmt.Println("> 4 node")
  longList4 := createListOfNumbers(1234)
  left4, right4 := frontBackSplit(longList4)
  printList(left4)
  fmt.Println(">>>")
  printList(right4)
  fmt.Println("")

  fmt.Println("> 9 nodes")
  longList5 := createListOfNumbers(123456789)
  left5, right5 := frontBackSplit(longList5)
  printList(left5)
  fmt.Println(">>>")
  printList(right5)
  fmt.Println("")

  fmt.Println("> nil check")
  left6, right6 := frontBackSplit(nil)
  printList(left6)
  fmt.Println(">>>")
  printList(right6)
  fmt.Println("")

}