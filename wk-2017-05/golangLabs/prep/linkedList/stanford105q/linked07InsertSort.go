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

// it should return the new head node.
// I think this was wrong.
// You were not doing some cleanup on the popped node.
// and you are not returning the thing you pop.
func pop(head *Node) (*Node, *Node){
  // sanity checks
  if head == nil {
    return nil, nil
  }

  // 1 node in system.
  if (*head).p == nil {
    return head, nil
  }

  // general case..
  tempN := (*head).p
  fmt.Println("deleting the head node", tempN)

  // setup return.
  firstP := head
  restP := (*head).p
  // nil the first node pointer since you are disconnecting it from the list.
  (*firstP).p = nil

  return firstP, restP
}

// This makes the assumption that the list is sorted.
func sortedInsert(headP, insertP *Node) (*Node) {
  // pre checks.
  if headP == nil || insertP == nil {
    return nil
  }
  // probably should check for p values.

  // probably need this.
  prevP := headP
  insertV := (*insertP).v

  for curP := headP ; ; curP = (*curP).p {
    if (*curP).v > insertV && curP == headP {
      // insert is the first node.
      (*insertP).p = curP
      return insertP
    } else if (*curP).v > insertV {
      // insert in middle of list.
      (*prevP).p = insertP
      (*insertP).p = curP
      
      // You did your insert so you are done.
      break
    } else if (*curP).p == nil {
      // you hit the end of the list. done. not quite..
      // the item is the largest in the list?
      (*curP).p = insertP
      break
    }

    // update the previous pointer.
    // this should be one behind curP.
    prevP = curP
  }

  return headP
}

func insertSort(headP *Node) (*Node) {
  // edge case to check.
  if headP == nil {
    return nil
  } else if (*headP).p == nil {
    return headP
  }

  // initial setup
  var finalP *Node
  finalP = nil

  // two nodes...
  for curP, restP := headP, headP ; curP != nil ; {
    // pop off a node of the existing list.
    curP, restP = pop(restP)

    if finalP == nil {
      // First node.
      finalP = curP
    } else if finalP != nil && curP != nil {
      // other node
      finalP = sortedInsert(finalP, curP)
      // fmt.Println(">>>>")
      // printList(finalP)
    }
  }

  // fmt.Println("<<<")
  // printList(finalP)
  return finalP
}



func main() {
  fmt.Println("> Insert sort test")
  longList1 := createListOfNumbers(54321)
  result1 := insertSort(longList1)
  printList(result1)

  fmt.Println("> Sorted")
  longList2 := createListOfNumbers(12345)
  result2 := insertSort(longList2)
  printList(result2)

  fmt.Println("> 1 item")
  longList3 := createListOfNumbers(1)
  result3 := insertSort(longList3)
  printList(result3)

  fmt.Println("> nil check")
  result4 := insertSort(nil)
  printList(result4)

  fmt.Println("> with duplicated and out of order")
  longList5 := createListOfNumbers(1302392)
  result5 := insertSort(longList5)
  printList(result5)
}