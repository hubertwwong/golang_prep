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



func main() {
  fmt.Println("> Sorted insert in middle of list")
  longList1 := createListOfNumbers(1235)
  myNode1 := Node{nil, 4}
  result1 := sortedInsert(longList1, &myNode1)
  printList(result1)

  fmt.Println("> Sorted insert in middle of list. Same number.")
  longList2 := createListOfNumbers(1235)
  myNode2 := Node{nil, 3}
  result2 := sortedInsert(longList2, &myNode2)
  printList(result2)

  fmt.Println("> Sorted insert in start of the list")
  longList3 := createListOfNumbers(1235)
  myNode3 := Node{nil, 9}
  result3 := sortedInsert(longList3, &myNode3)
  printList(result3)

  fmt.Println("> 1 item")
  longList4 := createListOfNumbers(5)
  myNode4 := Node{nil, 1}
  result4 := sortedInsert(longList4, &myNode4)
  printList(result4)

  fmt.Println("> 2 item")
  longList5 := createListOfNumbers(1)
  myNode5 := Node{nil, 5}
  result5 := sortedInsert(longList5, &myNode5)
  printList(result5)

  fmt.Println("> nil")
  result6 := sortedInsert(nil, nil)
  printList(result6)
}