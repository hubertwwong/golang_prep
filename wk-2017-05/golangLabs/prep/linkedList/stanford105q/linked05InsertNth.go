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

// nthItem is on array index. so it starts on 0th poisition
func insertNth(head *Node, nodeToInsert *Node, nthItem int) (*Node) {
  // pre checks.
  if head == nil {
    return nil
  } else if nodeToInsert == nil {
    return nil
  } else if nthItem < 0 {
    return nil
  }

  // initial variables.
  curIndex := 0
  prevP := head

  for curP := head ; ; curP = (*curP).p {
    if curIndex == nthItem && curIndex == 0 {
      (*nodeToInsert).p = head
      return nodeToInsert
    } else if curIndex == nthItem {
      // previous == new
      (*prevP).p = nodeToInsert

      // new.next == curPointer.
      (*nodeToInsert).p = curP
      
      // You did your insert so you are done.
      break
    } else if (*curP).p == nil {
      // end if the list.
      break
    }

    // update the previous pointer.
    // this should be one behind curP.
    prevP = curP

    // increment.
    curIndex++
  }

  return head
}



func main() {
  fmt.Println("> Insert in middle of list")
  longList := createListOfNumbers(12345)
  myNode := Node{nil, 0}
  result := insertNth(longList, &myNode, 3)
  printList(result)

  fmt.Println("> Insert at the start of list")
  longList1 := createListOfNumbers(12345)
  myNode1 := Node{nil, 0}
  result1 := insertNth(longList1, &myNode1, 0)
  printList(result1)

  fmt.Println("> Insert at the end of list")
  longList2 := createListOfNumbers(12345)
  myNode2 := Node{nil, 0}
  result2 := insertNth(longList2, &myNode2, 4)
  printList(result2)

  fmt.Println("> Insert single num")
  longList3 := createListOfNumbers(1)
  myNode3 := Node{nil, 0}
  result3 := insertNth(longList3, &myNode3, 0)
  printList(result3)

  fmt.Println("> Insert nil")
  myNode4 := Node{nil, 0}
  result4 := insertNth(nil, &myNode4, 0)
  printList(result4)
}